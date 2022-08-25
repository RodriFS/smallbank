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

	owner := initializers.DB.First(&models.User{}, body.UserId)
	if owner.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	account := models.Account{Owner: body.UserId}
	newAccount := initializers.DB.Create(&account)

	if newAccount.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating Account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}
