package datasources

import (
	"errors"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateAccount(account models.Account, db *gorm.DB) (models.Account, error) {
	result := db.First(&models.User{}, account.UserID)

	if result.Error != nil {
		return models.Account{}, errors.New("User: " + result.Error.Error())
	}

	result = db.Create(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func FindAccounts(db *gorm.DB) ([]models.Account, error) {
	var accounts []models.Account
	result := db.Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func FirstAccount(id string, db *gorm.DB) (models.Account, error) {
	var account models.Account
	result := db.First(&account, id)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func UpdateAccount(id string, account map[string]any, db *gorm.DB) error {
	existingAccount := models.Account{}
	result := db.First(&existingAccount, id)

	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existingAccount).Updates(account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteAccount(id string, db *gorm.DB) error {
	result := db.Delete(&models.Account{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
