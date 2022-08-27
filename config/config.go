package config

import "os"

var DB_HOST string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_PORT string

func DbCredentials() map[string]string {
	creds := make(map[string]string)
	creds["host"] = os.Getenv("DB_HOST")
	creds["user"] = os.Getenv("DB_USER")
	creds["password"] = os.Getenv("DB_PASSWORD")
	creds["name"] = os.Getenv("DB_NAME")
	creds["port"] = os.Getenv("DB_PORT")
	return creds
}
