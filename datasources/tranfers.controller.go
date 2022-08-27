package datasources

import (
	"errors"
	"fmt"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateTransfer(transfer models.Transfer, db *gorm.DB) (models.Transfer, error) {
	for _, userId := range []uint{transfer.FromAccountId, transfer.ToAccountId} {
		result := db.First(&models.User{}, userId)
		if result.Error != nil {
			return models.Transfer{}, errors.New(fmt.Sprintf("User %d not found", userId))
		}
	}

	result := db.Create(&transfer)
	if result.Error != nil {
		return models.Transfer{}, errors.New("Error while creating transfer")
	}

	return transfer, nil
}

func FindTransfers(userId string, db *gorm.DB) ([]models.Transfer, error) {
	var transfers []models.Transfer
	result := db.Find(&transfers, userId)
	if result.Error != nil {
		return nil, errors.New("Error while retrieving transfer  list")
	}

	return transfers, nil
}
