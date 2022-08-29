package controllers

import (
	"net/http"
	"smallbank/server/datasources"
	"smallbank/server/models"
	"smallbank/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateTransfer(c *gin.Context) {
	db := utils.GetDB(c)

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

	transfer, err := datasources.CreateTransfer(models.Transfer{
		FromAccountId: body.From,
		ToAccountId:   body.To,
		Amount:        body.Amount,
		Currency:      body.Currency,
	}, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transfer": transfer,
	})
}

func GetTransferList(c *gin.Context) {
	db := utils.GetDB(c)

	accountId := c.Param("AccountId")

	transfers, err := datasources.FindTransfersByAccountId(accountId, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transfers": transfers})
}
