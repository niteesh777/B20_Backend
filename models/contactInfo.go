package models

import (
	"gorm.io/gorm"
)

type ContactInfo struct {
	gorm.Model

	Id        int `gorm:"primarykey"`
	Real_name string
	Name      string
	Email     string
	Nick      string
}
