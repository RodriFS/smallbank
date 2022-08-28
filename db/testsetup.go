package db

import (
	"fmt"
	"log"
	"smallbank/server/config"
	"smallbank/server/initializers"
	"smallbank/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
}

func SetupTestDB() (conn *gorm.DB) {
	cred := config.DbCredentials()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/testing?sslmode=disable",
		cred["user"], cred["password"], cred["host"], cred["port"])

	db, err := gorm.Open(postgres.Open(connString))

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	db.AutoMigrate(
		models.User{},
		models.Account{},
		models.Transaction{},
		models.Transfer{},
	)

	return db
}

func CleanUpTestDB(db *gorm.DB) {
	migrator := db.Migrator()

	tableList, err := migrator.GetTables()

	if err != nil {
		log.Fatal("Unable to get database tables")
	}

	for _, table := range tableList {
		migrator.DropTable(table)
	}
}
