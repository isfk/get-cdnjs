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
	ownFiles := map[string]string{}
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
	// 下载文件并上传 TODO
	for fileName, fileUrl := range allFiles {
		fmt.Printf("开始抓取文件: %s\n", fileUrl)
		key := fmt.Sprintf("%s/%s/%s/%s", config.Conf.FilePath, libraryName, version, fileName)
		err = pkg.Fetch(config.Conf.Bucket, key, fileUrl)
		if err != nil {
			fmt.Printf("抓取失败: %s\n", err)
		}
		ownFile := fmt.Sprintf("%s/%s", config.Conf.CdnDomain, key)
		ownFiles[fileName] = ownFile
		fmt.Printf("抓取成功: %s\n\n", ownFile)
	}
	fmt.Printf("%s 所有文件已抓取到七牛:\n", libraryName)
	for _, v := range ownFiles {
		fmt.Println(v)
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
