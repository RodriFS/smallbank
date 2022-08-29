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
			return errors.New("AccountID: " + err.Error())
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

func FindTransactionsByAccountId(accountId string, db *gorm.DB) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := db.Where("account_id = ?", accountId).Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func FirstTransaction(id string, db *gorm.DB) (models.Transaction, error) {
	var transaction models.Transaction
	result := db.First(&transaction, id)
	if result.Error != nil {
		return models.Transaction{}, result.Error
	}

	return transaction, nil
}
