package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"smallbank/main/constants"
	"smallbank/main/models"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Madrid",
		constants.DB_HOST,
		constants.DB_USER,
		constants.DB_PASSWORD,
		constants.DB_NAME,
		constants.DB_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.Account{})
	DB.AutoMigrate(models.Transaction{})
	DB.AutoMigrate(models.Transfer{})
}
