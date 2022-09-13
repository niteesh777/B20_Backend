package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserId   int `gorm:"primarykey"`
	User     ContactInfo
	Name     string
	Email    string
	Password string
}
