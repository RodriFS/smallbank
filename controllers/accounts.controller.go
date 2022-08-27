package controllers

import (
	"net/http"
	"smallbank/server/datasources"
	"smallbank/server/models"
	"smallbank/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	db := utils.GetDB(c)

	var body struct {
		UserId   uint `binding:"required"`
		Currency string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId is required"})
		return
	}

	account, err := datasources.CreateAccount(models.Account{UserID: body.UserId, Currency: body.Currency}, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func GetAccountList(c *gin.Context) {
	db := utils.GetDB(c)

	accounts, err := datasources.FindAccounts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func GetAccount(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	account, err := datasources.FirstAccount(id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

func UpdateAccount(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	var body struct {
		Active bool
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body can't be empty"})
		return
	}

	err := datasources.UpdateAccount(id, map[string]any{"Active": body.Active}, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteAccount(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	err := datasources.DeleteAccount(id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
