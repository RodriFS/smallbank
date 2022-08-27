package controllers

import (
	"fmt"
	"net/http"
	"smallbank/server/initializers"
	"smallbank/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransfer(c *gin.Context) {
	var body struct {
		From     uint  `binding:"required"`
		To       uint  `binding:"required"`
		Amount   int64 `binding:"required"`
		Currency string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId is required"})
		return
	}

	var result *gorm.DB
	for _, userId := range []uint{body.From, body.To} {
		result = initializers.DB.First(&models.User{}, userId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User %d not found", userId)})
			return
		}
	}

	transfer := models.Transfer{ToAccountId: body.To, FromAccountId: body.From, Amount: body.Amount, Currency: body.Currency}
	result = initializers.DB.Create(&transfer)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating transfer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transfer": transfer,
	})
}

func GetTransferList(c *gin.Context) {
	userId := c.Param("UserId")

	var transfers []models.Transfer
	result := initializers.DB.Find(&transfers, userId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving transfer  list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transfers": transfers})
}
