package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDatabaseLayer() gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@/blog_golang?charset=utf8&parseTime=true")
	db.DB()
	return db
}
