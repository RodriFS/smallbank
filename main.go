package main

import (
	"smallbank/server/constants"
	"smallbank/server/controllers"
	"smallbank/server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	constants.LoadConfig()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/account/:id", controllers.GetAccount)
	r.GET("/account", controllers.GetAccountList)
	r.POST("/account", controllers.CreateAccount)
	r.PUT("/account/:id", controllers.UpdateAccount)
	r.DELETE("/account/:id", controllers.DeleteAccount)

	r.GET("/user/:id", controllers.GetUser)
	r.GET("/user", controllers.GetUserList)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}
