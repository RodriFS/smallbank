package controllers

import (
	"net/http"
	"smallbank/server/datasources"
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

	user, err := datasources.CreateUser(models.User{
		Name: body.Name,
		Last: body.Last,
		Phone: models.Phone{
			Code:   body.Phone.Code,
			Number: body.Phone.Number,
		},
	}, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUserList(c *gin.Context) {
	db := utils.GetDB(c)

	users, err := datasources.FindUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	user, err := datasources.FirstUser(id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	var Phone map[string]any
	if body.Phone != nil {
		Phone = map[string]any{
			"Code":   body.Phone.Code,
			"Number": body.Phone.Number,
		}
	}

	err := datasources.UpdateUser(id, map[string]any{
		"Name":  body.Name,
		"Last":  body.Last,
		"Phone": Phone,
	}, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	db := utils.GetDB(c)

	id := c.Param("id")

	err := datasources.DeleteUser(id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
