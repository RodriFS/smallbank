package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID       uint   `gorm:"index;not null"`
	Balance      int64  `gorm:"not null;default:0"`
	Currency     string `gorm:"not null;default:EUR"`
	Active       bool   `gorm:"not null;default:true"`
	Transactions []Transaction
	TransferTo   []Transfer `gorm:"foreignKey:ToAccountId"`
	TransferFrom []Transfer `gorm:"foreignKey:FromAccountId"`
}
