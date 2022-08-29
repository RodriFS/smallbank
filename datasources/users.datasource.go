package datasources

import (
	"errors"
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateUser(user models.User, db *gorm.DB) (models.User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, errors.New("Error while creating user")
	}

	return user, nil
}

func FindUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		return nil, errors.New("Error while retrieving user list")
	}

	return users, nil
}

func FirstUser(id string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Preload("Accounts").First(&user, id)

	if result.Error != nil {
		return models.User{}, errors.New("Error while retrieving user")
	}

	return user, nil
}

func UpdateUser(id string, user map[string]any, db *gorm.DB) error {
	var existingUser models.User
	result := db.First(&existingUser, id)

	if result.Error != nil {
		return errors.New("Error while retrieving user")
	}

	result = db.Model(&existingUser).Updates(&user)
	if result.Error != nil {
		return errors.New("Error while retrieving user")
	}

	return nil
}

func DeleteUser(id string, db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", id).Delete(&models.Account{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.User{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
