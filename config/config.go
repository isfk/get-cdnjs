package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	CdnDomain string `mapstructure:"cdn_domain"`
	Proxy     string `mapstructure:"proxy"`
	FilePath  string `mapstructure:"file_path"`
}

var Conf = &Config{}

func InitConfig() {
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	if !strings.Contains(Conf.Proxy, "http") {
		fmt.Println("当前配置不使用代理")
	}
}
