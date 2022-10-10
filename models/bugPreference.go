package models

import (
	"gorm.io/gorm"
)

type BugPreference struct {
	gorm.Model

	UserID int `gorm:"primarykey"`
	User   ContactInfo

	Comment_count      bool
	Type               bool
	Status             bool
	Priority           bool
	Severity           bool
	Summary            bool
	Product            bool
	Platform           bool
	Resolution         bool
	Target_milestone   bool
	Classification     bool
	Is_confirmed       bool
	Is_open            bool
	Qa_contact         bool
	Creator_detail     bool
	Assigned_to_detail bool
	Last_change_time   bool
	Creation_time      bool
}
