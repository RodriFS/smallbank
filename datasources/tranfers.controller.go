package datasources

import (
	"errors"
	"fmt"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateTransfer(transfer models.Transfer, db *gorm.DB) (models.Transfer, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, userId := range []uint{transfer.FromAccountId, transfer.ToAccountId} {
			if err := tx.First(&models.User{}, userId).Error; err != nil {
				return errors.New(fmt.Sprintf("User %d: %q", userId, err))
			}
		}

		if err := tx.Create(&transfer).Error; err != nil {
			return err
		}

		var receivingAccount models.Account
		if err := tx.First(&receivingAccount, transfer.ToAccountId).Error; err != nil {
			return err
		}

		receivingAccount.Balance += transfer.Amount
		if err := tx.Save(&receivingAccount).Error; err != nil {
			return err
		}

		return nil
	})

	return transfer, err
}

func FindTransfers(userId string, db *gorm.DB) ([]models.Transfer, error) {
	var transfers []models.Transfer
	result := db.Find(&transfers, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return transfers, nil
}
