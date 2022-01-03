package main

import (
	"imask-go/controller"
	"imask-go/model"
)

type Users struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
}

func main() {
	model.InitDB()
	controller.InitEcho()
}
