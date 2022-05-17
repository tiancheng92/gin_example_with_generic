package mysql

import (
	"gin_example_with_generic/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetGormClient(dsn string, tables []any) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("%+v", err)
	}
	return db
}
