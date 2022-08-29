package models

import (
	"errors"

	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ToAccountId   uint   `gorm:"index;index:transfer;not null"`
	FromAccountId uint   `gorm:"index;index:transfer;not null"`
	Amount        int64  `gorm:"not null"`
	Currency      string `gorm:"not null"`
}

func (t *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Amount <= 0 {
		err = errors.New("Can't transfer 0 or less amount")
	}

	if t.FromAccountId == t.ToAccountId {
		err = errors.New("A user can't make a transfer to itself")
	}

	return
}
