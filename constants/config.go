package constants

import "os"

var DB_HOST string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_PORT string

func LoadConfig() {
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
}
