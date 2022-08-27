package datasources

import (
	"errors"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateTransaction(transaction models.Transaction, db *gorm.DB) (models.Transaction, error) {
	result := db.First(&models.Account{}, transaction.AccountID)
	if result.Error != nil {
		return models.Transaction{}, errors.New("Account not found")
	}

	result = db.Create(&transaction)
	if result.Error != nil {
		return models.Transaction{}, errors.New("Error while creating transaction")
	}

	return transaction, nil
}

func FindTransactions(userId string, db *gorm.DB) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := db.Find(&transactions, userId)

	if result.Error != nil {
		return nil, errors.New("Error while retrieving transactions list")
	}

	return transactions, nil
}
