package qiniu

import (
	"fmt"
	"os"
	"strings"

	"github.com/isfk/get-cdnjs/config"
	"github.com/isfk/get-cdnjs/pkg"
)

func RunList() {
	config.InitConfig()
	var err error
	list := []string{}
	childList := []string{}
	result := map[string][]string{}
	if config.Conf.FilePath != "" {
		list, err = pkg.List(config.Conf.Bucket, fmt.Sprintf("%s/", config.Conf.FilePath))
	} else {
		list, err = pkg.List(config.Conf.Bucket, "")
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range list {
		if strings.HasSuffix(v, "/") {
			childList, err = pkg.List(config.Conf.Bucket, v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			result[v] = childList
		}
	}
	fmt.Printf("─ %s:\n", config.Conf.FilePath)
	for p, v := range result {
		fmt.Printf("  ├─ %s:\n", strings.Trim(strings.ReplaceAll(p, config.Conf.FilePath, ""), "/"))
		for _, c := range v {
			fmt.Printf("  │  └─ %s\n", strings.Trim(strings.ReplaceAll(c, p, ""), "/"))
		}
	}
}
