package entity

import (
	"gorm.io/gorm"
)
type TypeFund struct {
	gorm.Model
	TypeFund string
	Donators []Donator `grom:"foreignKey:TypeFundID"`
}

type Organization struct {
	gorm.Model
	Organization string
	Donators []Donator `grom:"foreignKey:OrganizationID"`
}

type Donator struct {
	gorm.Model

	AdminID *uint
	Admin   Admin

	UserName string
	DateTime string
	UserInfo string
	UserNote string
	Amount   int
	NameFund string

	TypeFundID *uint
	TypeFund   TypeFund
	OrganizationID *uint
	Organization   Organization
}