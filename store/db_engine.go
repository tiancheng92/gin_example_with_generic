package store

import (
	"gin_example_with_generic/config"
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

func Init() {
	initDefaultDB()
}
