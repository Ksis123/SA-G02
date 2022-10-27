package entity

import (
	"gorm.io/gorm"
)

//ระบบลงทะเบียนข้อมูลนักศึกษา---------------------------------------------------------------
type Year struct {
	gorm.Model

	Number string

	Students []Student `grom:"foreignKey:YearId"`
}

type Faculty struct {
	gorm.Model

	Name     string
	ThaiName string

	Students []Student `grom:"foreignKey:FacultyId"`
}

type Advisor struct {
	gorm.Model

	Name     string
	ThaiName string

	Students []Student `grom:"foreignKey:AdvisorId"`
}

type Student struct {
	gorm.Model

	Personalid string
	Name       string
	Phon       string
	Gpax       float64
	Money      int

	YearId *uint
	Year   Year

	FacultyId *uint
	Faculty   Faculty

	AdvisorId *uint
	Advisor   Advisor

	UserId *uint `gorm:"uniqueIndex"`
	User   User

	Reports []Report `grom:"foreignKey:StudentId"`
}
