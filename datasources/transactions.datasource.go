package datasources

import (
	"errors"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateTransaction(transaction models.Transaction, db *gorm.DB) (models.Transaction, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.First(&account, transaction.AccountID).Error; err != nil {
			return err
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		account.Balance += transaction.Amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		return nil
	})

	return transaction, err
}

func FindTransactions(userId string, db *gorm.DB) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := db.Find(&transactions, userId)

	if result.Error != nil {
		return nil, errors.New("Error while retrieving transactions list")
	}

	return transactions, nil
}
