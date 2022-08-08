package config

import (
	"gin_example_with_generic/types/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	configInfo            *config.Config
	HotUpdateForStone     = make(chan struct{})
	HotUpdateForLog       = make(chan struct{})
	HotUpdateForValidator = make(chan struct{})
	HotUpdateForServer    = make(chan struct{})
)

func Init() {
	viper.SetConfigFile("./config_file/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&configInfo); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		oldCfg := *configInfo
		if err := viper.Unmarshal(&configInfo); err != nil {
			panic(err)
		}
		newCfg := *configInfo
		if oldCfg.Server.ServicePort != newCfg.Server.ServicePort || oldCfg.Server.ServiceHost != newCfg.Server.ServiceHost || oldCfg.Server.Mode != newCfg.Server.Mode {
			HotUpdateForServer <- struct{}{}
		}
		if oldCfg.Mysql != newCfg.Mysql {
			HotUpdateForStone <- struct{}{}
		}
		if oldCfg.LogLevel != newCfg.LogLevel {
			HotUpdateForLog <- struct{}{}
		}
		if oldCfg.I18n != newCfg.I18n {
			HotUpdateForValidator <- struct{}{}
		}
	})
}

func GetConf() *config.Config {
	return configInfo
}
