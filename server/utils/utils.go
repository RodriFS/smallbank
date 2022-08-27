package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDB(c *gin.Context) *gorm.DB {
	db := c.Value("DB")
	if db == nil {
		log.Fatal("Database not attached")
	}

	return db.(*gorm.DB)
}
