package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "imask_dev:imask_dev@tcp(127.0.0.1:3306)/imask_dev?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&User{}, &Media{}, &Task{}, &Annotation{})
	if err != nil {
		panic(err)
	}
}
