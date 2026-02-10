package qiniu

import (
	"fmt"
	"os"

	"github.com/isfk/get-cdnjs/config"
	"github.com/isfk/get-cdnjs/pkg"
)

func RunDelete(targetPath string) {
	config.InitConfig()

	pathToDelete := targetPath
	if pathToDelete == "" {
		fmt.Println("请输入要删除的目录:")
		_, err := fmt.Scanln(&pathToDelete)
		if err != nil {
			fmt.Println("输入错误")
			os.Exit(1)
		}
	}

	if pathToDelete == "" {
		fmt.Println("错误: 目录路径不能为空")
		os.Exit(1)
	}

	fullPath := pathToDelete
	if config.Conf.FilePath != "" {
		fullPath = fmt.Sprintf("%s/%s", config.Conf.FilePath, pathToDelete)
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

	fmt.Printf("找到以下目录:\n")
	for _, v := range list {
		fmt.Printf("  - %s\n", v)
	}

	fmt.Print("\n确认删除以上目录及其所有文件? (yes/no): ")
	var confirm string
	_, err = fmt.Scanln(&confirm)
	if err != nil || confirm != "yes" {
		fmt.Println("取消删除")
		os.Exit(0)
	}

	count, err := pkg.Delete(config.Conf.Bucket, fullPath+"/")
	if err != nil {
		fmt.Printf("删除失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n成功删除 %d 个文件\n", count)
	fmt.Printf("目录: %s\n", fullPath)
}
