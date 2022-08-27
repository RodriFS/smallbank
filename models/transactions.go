package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	AccountID uint   `gorm:"index;not null"`
	Amount    int64  `gorm:"not null"`
	Currency  string `gorm:"not null"`
}
