package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func LoadEnvVariables() {
	err := godotenv.Load(RootDir() + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file, " + err.Error())
	}
}
