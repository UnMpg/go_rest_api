package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dtb := "root:@tcp(127.0.0.1:3306)/go_rest_api"
	db, err := gorm.Open(mysql.Open(dtb), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
