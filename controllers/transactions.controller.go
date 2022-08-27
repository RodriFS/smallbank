package controllers

import (
	"net/http"
	"smallbank/server/datasources"
	"smallbank/server/models"
	"smallbank/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	db := utils.GetDB(c)

	var body struct {
		AccountId uint  `binding:"required"`
		Amount    int64 `binding:"required"`
		Currency  string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId is required"})
		return
	}

	transaction, err := datasources.CreateTransaction(
		models.Transaction{
			AccountID: body.AccountId,
			Amount:    body.Amount,
			Currency:  body.Currency,
		}, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}

func GetTransactionList(c *gin.Context) {
	db := utils.GetDB(c)

	userId := c.Param("UserId")

	transactions, err := datasources.FindTransactions(userId, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
