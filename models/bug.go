package models

import (
	"time"

	"gorm.io/gorm"
)

type Bug struct {
	gorm.Model

	Id            float64 `gorm:"primarykey"`
	Comment_count float64
	// Deadline           string
	Type                 string
	Status               string
	Priority             string
	Severity             string
	Summary              string
	Product              string
	Platform             string
	Resolution           string
	Target_milestone     string
	Classification       string
	Is_confirmed         bool
	Is_open              bool
	Qa_contactID         int
	Creator_detailID     int
	Assigned_to_detailID int
	Qa_contact           ContactInfo
	Creator_detail       ContactInfo
	Assigned_to_detail   ContactInfo
	//Cc_details         []ContactInfo `gorm:"foreignkey:Id"`
	Last_change_time time.Time
	Creation_time    time.Time

	// Assigned_to        string
	// Creator            string
}
