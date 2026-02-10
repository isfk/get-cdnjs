package qiniu

import (
	"fmt"
	"os"
	"strings"

	"github.com/isfk/get-cdnjs/config"
	"github.com/isfk/get-cdnjs/pkg"
)

func RunDelete(targetPath string) {
	config.InitConfig()

	if targetPath == "" {
		fmt.Println("\033[36m=== 七牛云现有目录 ===\033[0m")

		var list []string
		var err error
		if config.Conf.FilePath != "" {
			list, err = pkg.List(config.Conf.Bucket, fmt.Sprintf("%s/", config.Conf.FilePath))
		} else {
			list, err = pkg.List(config.Conf.Bucket, "")
		}
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			os.Exit(1)
		}

		if len(list) == 0 {
			fmt.Println("暂无目录")
			os.Exit(0)
		}

		for _, v := range list {
			if strings.HasSuffix(v, "/") {
				relativePath := strings.Trim(strings.ReplaceAll(v, config.Conf.FilePath, ""), "/")
				fmt.Printf("  - %s\n", strings.Trim(relativePath, "/"))
			}
		}

		fmt.Println("\n请输入要删除的目录:")
		_, err = fmt.Scanln(&targetPath)
		if err != nil {
			fmt.Println("输入错误")
			os.Exit(1)
		}
	}

	if targetPath == "" {
		fmt.Println("错误: 目录路径不能为空")
		os.Exit(1)
	}

	fullPath := targetPath
	if config.Conf.FilePath != "" {
		fullPath = fmt.Sprintf("%s/%s", config.Conf.FilePath, targetPath)
	}

	list, err := pkg.List(config.Conf.Bucket, fullPath+"/")
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
		os.Exit(1)
	}

	if len(list) == 0 {
		fmt.Printf("目录不存在或为空: %s\n", fullPath)
		os.Exit(0)
	}

	result := map[string][]string{}
	for _, dir := range list {
		if strings.HasSuffix(dir, "/") {
			files, err := pkg.List(config.Conf.Bucket, dir)
			if err != nil {
				fmt.Printf("查询文件失败 %s: %v\n", dir, err)
				continue
			}
			result[dir] = files
		}
	}

	fmt.Printf("\033[31m=== 即将删除以下目录和文件 ===\033[0m\n")
	fmt.Printf("路径: %s\n\n", fullPath)

	totalFiles := 0
	for dir, files := range result {
		relativePath := strings.Trim(strings.ReplaceAll(dir, fullPath, ""), "/")
		fmt.Printf("  \033[33m├─ %s/:\033[0m\n", relativePath)
		for _, f := range files {
			fileName := strings.Trim(strings.ReplaceAll(f, dir, ""), "/")
			if fileName != "" {
				fmt.Printf("  │  └─ %s\n", fileName)
				totalFiles++
			}
		}
	}

	fmt.Printf("\n共 %d 个目录，%d 个文件\n", len(result), totalFiles)
	fmt.Print("\n\033[31m确认删除? (输入 'yes' 继续):\033[0m ")
	var confirm string
	_, err = fmt.Scanln(&confirm)
	if err != nil || confirm != "yes" {
		fmt.Println("取消删除")
		os.Exit(0)
	}

	count, err := pkg.Delete(config.Conf.Bucket, fullPath+"/")
	if err != nil {
		fmt.Printf("\n删除失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\033[32m✓ 成功删除 %d 个文件\033[0m\n", count)
	fmt.Printf("目录: %s\n", fullPath)
}
