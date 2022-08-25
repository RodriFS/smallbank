package controllers

import (
	"net/http"
	"smallbank/main/initializers"
	"smallbank/main/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	type phone struct {
		Code   int   `binding:"required"`
		Number int32 `binding:"required"`
	}
	var body struct {
		Name  string `binding:"required"`
		Last  string `binding:"required"`
		Phone *phone `binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := models.User{
		Name: body.Name,
		Last: body.Last,
		Phone: models.Phone{
			Code:   body.Phone.Code,
			Number: body.Phone.Number,
		},
	}
	newAccount := initializers.DB.Create(&user)

	if newAccount.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating Account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
