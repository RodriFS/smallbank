package models

import (
	"errors"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID uint   `gorm:"index;not null"`
	Amount    int64  `gorm:"not null"`
	Currency  string `gorm:"not null"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Amount == 0 {
		err = errors.New("Can't do a transaction on 0 value")
	}

	return
}
