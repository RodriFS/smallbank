package datasources

import (
	"smallbank/server/models"

	"gorm.io/gorm"
)

func CreateUser(user models.User, db *gorm.DB) (models.User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func FindUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func FirstUser(id string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Preload("Accounts").First(&user, id)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func UpdateUser(id string, user map[string]any, db *gorm.DB) error {
	var existingUser models.User
	result := db.First(&existingUser, id)

	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existingUser).Updates(&user)
	if result.Error != nil {
		return result.Error
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
