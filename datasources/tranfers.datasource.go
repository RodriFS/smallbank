package datasources

import (
	"errors"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateTransfer(transfer models.Transfer, db *gorm.DB) (models.Transfer, error) {
	err := db.Transaction(func(tx *gorm.DB) error {

		var fromAccount models.Account
		if err := tx.First(&fromAccount, transfer.FromAccountId).Error; err != nil {
			return errors.New("FromAccountId: " + err.Error())
		}

		fromAccount.Balance -= transfer.Amount

		if fromAccount.Balance < 0 {
			return errors.New("FromAccountId: Insufficient funds")
		}

		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}

		var toAccount models.Account
		if err := tx.First(&toAccount, transfer.ToAccountId).Error; err != nil {
			return errors.New("ToAccountId: " + err.Error())
		}

		toAccount.Balance += transfer.Amount
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}

		if err := tx.Create(&transfer).Error; err != nil {
			return err
		}

		return nil
	})

	return transfer, err
}

func FindTransfers(accountId string, db *gorm.DB) ([]models.Transfer, error) {
	var transfers []models.Transfer
	result := db.Where("from_account_id = ?", accountId).Or("to_account_id = ?", accountId).Find(&transfers)
	if result.Error != nil {
		return nil, result.Error
	}

	return transfers, nil
}

func FirstTransfer(id string, db *gorm.DB) (models.Transfer, error) {
	var transfer models.Transfer
	result := db.First(&transfer, id)
	if result.Error != nil {
		return models.Transfer{}, result.Error
	}

	return transfer, nil
}
