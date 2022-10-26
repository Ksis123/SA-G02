package entity

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name          string
	Email         string        `gorm:"uniqueIndex"`
	Password      string        `json:"-"`
	SlipLists     []SlipList    `gorm:"foreignKey:AdminID"`
	Donators      []Donator     `gorm:"foreignKey:AdminID"`
	Scholarshipes []Scholarship `gorm:"foreignKey:AdminID"`
	StudentLists  []StudentList `gorm:"foreignKey:AdminID"`
}
