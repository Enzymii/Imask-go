package model

type User struct {
	Username string `gorm:"primary_key;size:30"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null,unique"`
}
