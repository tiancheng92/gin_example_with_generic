package config

import (
	"gin_example_with_generic/types/config"
	"github.com/spf13/viper"
)

var (
	configInfo *config.Config
)

func Init(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&configInfo); err != nil {
		panic(err)
	}
}

func GetConf() *config.Config {
	return configInfo
}
