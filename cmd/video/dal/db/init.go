package db

import (
	"micro_tiktok/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormOpentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormOpentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&Video{}) {
		return
	}

	// 表不存在则在数据库中创建表
	if err = m.CreateTable(&Video{}); err != nil {
		panic(err)
	}
}
