package controllers

import (
	"net/http"
	"smallbank/server/initializers"
	"smallbank/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	var body struct {
		AccountId uint  `binding:"required"`
		Amount    int64 `binding:"required"`
		Currency  string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId is required"})
		return
	}

	var result *gorm.DB
	result = initializers.DB.First(&models.User{}, body.AccountId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	transaction := models.Transaction{AccountID: body.AccountId, Amount: body.Amount, Currency: body.Currency}
	result = initializers.DB.Create(&transaction)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}

func GetTransactionList(c *gin.Context) {
	var transactions []models.Transaction
	result := initializers.DB.Find(&transactions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving transactions list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
