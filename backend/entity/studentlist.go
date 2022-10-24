package entity

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Name         string
	StudentLists []StudentList `gorm:"foreignKey:ReportID"`
}
type Status struct {
	gorm.Model
	Status       string
	StudentLists []StudentList `gorm:"foreignKey:StatusID"`
}
type StudentList struct {
	gorm.Model
	Reason string
	Amount int

	ReportID *int
	Report   Report `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`
	// Status    string
	// Report	  string
	SaveTime  time.Time
	SlipLists []SlipList `gorm:"foreignKey:StudentListID"`
}
