package controllers

import (
	"net/http"
	"smallbank/server/models"
	"smallbank/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := utils.GetDB(c)

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
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUserList(c *gin.Context) {
	db := utils.GetDB(c)

	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	var user models.User
	result := db.Preload("Accounts").First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	db := utils.GetDB(c)

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
	result := db.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving user"})
		return
	}

	var Phone models.Phone
	if body.Phone != nil {
		Phone = models.Phone{
			Code:   body.Phone.Code,
			Number: body.Phone.Number,
		}
	}

	result = db.Model(&user).Updates(models.User{Name: body.Name, Last: body.Last, Phone: Phone})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while retrieving user"})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	result := db.Preload("Accounts").Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while deleting user"})
		return
	}

	c.Status(http.StatusOK)
}
