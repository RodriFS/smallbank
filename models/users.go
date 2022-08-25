package models

import "gorm.io/gorm"

type Phone struct {
	Code   int
	Number int32
}

type User struct {
	gorm.Model
	Name    string    `gorm:"not null"`
	Last    string    `gorm:"index;not null"`
	Phone   Phone     `gorm:"embedded;unique;not null"`
	Account []Account `gorm:"foreignKey:Owner"`
}
