package controllers

import (
	"net/http"
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

	result := db.First(&models.User{}, body.UserId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	account := models.Account{UserID: body.UserId}
	result = db.Create(&account)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func GetAccountList(c *gin.Context) {
	db := utils.GetDB(c)

	var accounts []models.Account
	result := db.Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func GetAccount(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	var account models.Account
	result := db.First(&account, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account"})
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

	var account models.Account
	result := db.First(&account, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account"})
		return
	}

	result = db.Model(&account).Updates(map[string]interface{}{"active": body.Active})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while retrieving account"})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteAccount(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	result := db.Delete(&models.Account{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while deleting account"})
		return
	}

	c.Status(http.StatusOK)
}
