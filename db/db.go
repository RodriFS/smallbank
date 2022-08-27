package db

import (
	"fmt"
	"log"
	"smallbank/server/config"
	"smallbank/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (conn *gorm.DB) {
	cred := config.DbCredentials()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cred["user"], cred["password"], cred["host"], cred["port"], cred["name"])

	db = connectPool("main", connString)

	db.AutoMigrate(
		models.User{},
		models.Account{},
		models.Transaction{},
		models.Transfer{},
	)
	return db
}

func connectPool(applicationName string, connString string) (conn *gorm.DB) {
	db, err := gorm.Open(postgres.Open(connString))

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	log.Println("Successfully connected")

	return db
}
