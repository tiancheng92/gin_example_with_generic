package store

import (
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/mysql"
	"gin_example_with_generic/store/model"
	"github.com/tiancheng92/gf"
	"gorm.io/gorm"
)

var (
	defaultDB       *gorm.DB
	defaultDBTables = []any{new(model.Country), new(model.User)}
)

func GetDefaultDB() *gorm.DB {
	if config.GetConf().Log.Level == "debug" {
		return defaultDB.Debug()
	}
	return defaultDB
}

func InitDefaultDB() {
	var (
		conf = config.GetConf()
		dsn  = gf.StringJoin(conf.Mysql.DBUser, ":", conf.Mysql.DBPassword, "@tcp(", conf.Mysql.DBHost, ")/", conf.Mysql.DBName, "?charset=utf8&parseTime=true&loc=Local")
	)
	defaultDB = mysql.GetGormClient(dsn, defaultDBTables)
}
