package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormOpentracing "gorm.io/plugin/opentracing"
	"micro_tiktok/pkg/constants"
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

	if m.HasTable(&Comment{}) && m.HasTable(&CommentCount{}) {
		return
	}

	// 表不存在则在数据库中创建表
	err = m.DropTable(&Comment{}, &CommentCount{})
	if err != nil {
		panic(err)
	}
	if err = m.AutoMigrate(&Comment{}, &CommentCount{}); err != nil {
		panic(err)
	}
}
