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
		Name:  body.Name,
		Last:  body.Last,
		Phone: models.Phone(*body.Phone),
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

func GetUserList(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := initializers.DB.Preload("Accounts").First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	type phone struct {
		Code   int
		Number int32
	}

	var body struct {
		Name  string
		Last  string
		Phone *phone
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body can't be empty"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user"})
		return
	}

	result = initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Last: body.Last, Phone: models.Phone(*body.Phone)})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while retrieving user"})
		return
	}

	c.Status(http.StatusOK)
}
