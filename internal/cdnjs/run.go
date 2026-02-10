package cdnjs

import (
	"fmt"
	"os"

	"github.com/isfk/get-cdnjs/config"
	"github.com/isfk/get-cdnjs/pkg"
)

func Run() {
	config.InitConfig()
	for {
		Scan()
	}
}

func Scan() {
	var libraryName string
	var version string
	var confirm string

	fmt.Print("请输入库名: ")
	_, err := fmt.Scan(&libraryName)
	if err != nil {
		fmt.Println("输入错误")
		os.Exit(1)
	}
	fmt.Println("正在查找版本号, 请稍等...")

	// 列出所有版本号
	version, versions := GetVersions(libraryName)
	if version == "" {
		fmt.Printf("未找到库或查询超时: %s\n", libraryName)
		Scan()
		return
	}
	fmt.Printf("最新版本号: \033[31m%s\033[0m, 可选版本号: \n", version)
	fmt.Print("\033[31m[ \033[0m")
	for _, v := range versions {
		fmt.Printf("%s ", v)
	}
	fmt.Print("\033[31m]\033[0m")
	fmt.Println()
	fmt.Print("\033[31m回车使用最新版本号\033[0m, 否则请输入版本号: ")
	_, _ = fmt.Scanln(&version)
	fmt.Printf("库名: \033[31m%s\033[0m, 版本号: \033[31m%s\033[0m\n", libraryName, version)
	// 列出所有文件
	fmt.Println("正在查找文件, 请稍等...")
	files := GetFiles(libraryName, version)
	fmt.Println("该版本存在以下文件: ")

	allFiles := map[string]string{}
	fmt.Print("\033[31m[ \033[0m")
	for _, v := range files {
		fmt.Printf("%s ", v)
		allFiles[v] = fmt.Sprintf("https://cdn.jsdelivr.net/npm/%s@%s/%s", libraryName, version, v)
	}
	fmt.Print("\033[31m]\033[0m")
	fmt.Println()
	fmt.Print("回车开始抓取, 结束请按键: <Ctrl + C> ")
	_, _ = fmt.Scanln(&confirm)
	fmt.Println()

	downloadedFiles := map[string][]byte{}
	downloadFailed := []struct {
		name string
		url  string
		err  error
	}{}

	fmt.Println("\033[36m=== 第一步：下载文件 ===\033[0m")
	for fileName, fileUrl := range allFiles {
		var data []byte
		var downloadErr error

		for retry := 0; retry < 3; retry++ {
			if retry > 0 {
				fmt.Printf("重试 %d/3: %s\n", retry, fileUrl)
			} else {
				fmt.Printf("下载: %s\n", fileUrl)
			}

			data, downloadErr = GetFileBytes(fileUrl, config.Conf.Proxy)
			if downloadErr == nil {
				break
			}
		}

		if downloadErr != nil {
			fmt.Printf("  \033[31m下载失败（已重试3次）: %v\033[0m\n\n", downloadErr)
			downloadFailed = append(downloadFailed, struct {
				name string
				url  string
				err  error
			}{fileName, fileUrl, downloadErr})
			continue
		}

		downloadedFiles[fileName] = data
		fmt.Printf("  \033[32m下载成功: %d bytes\033[0m\n\n", len(data))
	}

	ownFiles := map[string]string{}
	uploadFailed := []struct {
		name string
		err  error
	}{}

	if len(downloadedFiles) > 0 {
		fmt.Println("\n\033[36m=== 第二步：上传文件 ===\033[0m")
		for fileName, data := range downloadedFiles {
			key := fmt.Sprintf("%s/%s/%s/%s", config.Conf.FilePath, libraryName, version, fileName)

			uploadErr := pkg.Upload(key, data)
			if uploadErr != nil {
				fmt.Printf("上传失败: %s - \033[31m%v\033[0m\n\n", fileName, uploadErr)
				uploadFailed = append(uploadFailed, struct {
					name string
					err  error
				}{fileName, uploadErr})
				continue
			}

			ownFile := fmt.Sprintf("%s/%s", config.Conf.CdnDomain, key)
			ownFiles[fileName] = ownFile
			fmt.Printf("上传成功: %s\n\n", ownFile)
		}
	}

	fmt.Printf("\n\033[32m=== %s @ %s 抓取完成 ===\033[0m\n", libraryName, version)
	fmt.Printf("下载成功: %d 个文件\n", len(downloadedFiles))
	fmt.Printf("下载失败: %d 个文件\n", len(downloadFailed))
	fmt.Printf("上传成功: %d 个文件\n", len(ownFiles))
	fmt.Printf("上传失败: %d 个文件\n", len(uploadFailed))

	if len(downloadFailed) > 0 {
		fmt.Println("\n\033[31m下载失败文件:\033[0m")
		for _, f := range downloadFailed {
			fmt.Printf("  - %s\n", f.name)
			fmt.Printf("    URL: %s\n", f.url)
			fmt.Printf("    错误: %v\n", f.err)
		}
	}

	if len(uploadFailed) > 0 {
		fmt.Println("\n\033[31m上传失败文件:\033[0m")
		for _, f := range uploadFailed {
			fmt.Printf("  - %s\n", f.name)
			fmt.Printf("    错误: %v\n", f.err)
		}
	}

	if len(ownFiles) > 0 {
		fmt.Println("\n\033[32m成功上传:\033[0m")
		for _, v := range ownFiles {
			fmt.Println(v)
		}
	}

	fmt.Println()
	fmt.Println("结束请按键: <Ctrl + C> ")
	fmt.Println()
	fmt.Println()
}

type JSDelivrVersionsRet struct {
	Tags     Tags     `json:"tags"`
	Versions []string `json:"versions"`
}

type Tags struct {
	Latest string `json:"latest"`
}

type FileEntry struct {
	Type  string      `json:"type"` // "file" or "directory"
	Name  string      `json:"name"`
	Files []FileEntry `json:"files,omitempty"`
}

type JSDelivrFilesRet struct {
	Default string      `json:"default"`
	Files   []FileEntry `json:"files"`
}

// 保留旧结构体用于兼容
type VersionsRet struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Versions []string `json:"versions"`
}

type FilesRet struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Files   []string `json:"files"`
}

func flattenFiles(entries []FileEntry, prefix string) []string {
	var files []string
	for _, e := range entries {
		path := e.Name
		if prefix != "" {
			path = prefix + "/" + e.Name
		}
		if e.Type == "file" {
			files = append(files, path)
		} else if len(e.Files) > 0 {
			files = append(files, flattenFiles(e.Files, path)...)
		}
	}
	return files
}

func GetVersions(libraryName string) (version string, versions []string) {
	ret := &JSDelivrVersionsRet{}
	url := fmt.Sprintf("https://data.jsdelivr.com/v1/package/npm/%s", libraryName)
	fmt.Printf("查找链接: %s\n", url)
	_, err := Get[JSDelivrVersionsRet](url, config.Conf.Proxy, ret)
	if err != nil {
		return
	}

	return ret.Tags.Latest, ret.Versions
}

func GetFiles(libraryName string, version string) (files []string) {
	ret := &JSDelivrFilesRet{}
	url := fmt.Sprintf("https://data.jsdelivr.com/v1/package/npm/%s@%s", libraryName, version)
	fmt.Printf("查找链接: %s\n", url)
	_, err := Get[JSDelivrFilesRet](url, config.Conf.Proxy, ret)
	if err != nil {
		return
	}

	return flattenFiles(ret.Files, "")
}
