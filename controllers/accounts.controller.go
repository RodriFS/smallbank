package controllers

import (
	"net/http"
	"smallbank/main/initializers"
	"smallbank/main/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var body struct {
		UserId   uint `binding:"required"`
		Currency string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId is required"})
		return
	}

	result := initializers.DB.First(&models.User{}, body.UserId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	account := models.Account{UserID: body.UserId}
	result = initializers.DB.Create(&account)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func GetAccountList(c *gin.Context) {
	var accounts []models.Account
	result := initializers.DB.Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func GetAccount(c *gin.Context) {
	id := c.Param("id")

	var account models.Account
	result := initializers.DB.First(&account, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

func UpdateAccount(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Active bool
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body can't be empty"})
		return
	}

	var account models.Account
	result := initializers.DB.First(&account, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving account"})
		return
	}

	result = initializers.DB.Model(&account).Updates(map[string]interface{}{"active": body.Active})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while retrieving account"})
		return
	}

	c.Status(http.StatusOK)
}
