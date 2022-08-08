package store

import (
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/pkg/mysql"
	"gin_example_with_generic/store/model"
	"gorm.io/gorm"
)

var (
	defaultDB       *gorm.DB
	defaultDBTables = []any{new(model.Country), new(model.User)}
)

func GetDefaultDB() *gorm.DB {
	if config.GetConf().LogLevel == "debug" {
		return defaultDB.Debug()
	}
	return defaultDB
}

func initDefaultDB() {
	defaultDB = mysql.GetGormClient(config.GetConf().Mysql, defaultDBTables)
}

func initStore() {
	initDefaultDB()
}

func Init() {
	initStore()

	go func() {
		for {
			select {
			case <-config.HotUpdateForStone:
				initStore()
				log.Info("Store 热更新完成。")
			}
		}
	}()
}
