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
	r.GET("/accounts/:id", controllers.GetAccount)
	r.GET("/accounts", controllers.GetAccountList)
	r.POST("/accounts", controllers.CreateAccount)
	r.PUT("/accounts/:id", controllers.UpdateAccount)
	r.DELETE("/accounts/:id", controllers.DeleteAccount)

	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users", controllers.GetUserList)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/transfers/:UserId", controllers.GetTransferList)
	r.POST("/transfers", controllers.CreateTransfer)

	r.GET("/transactions/:UserId", controllers.GetTransactionList)
	r.POST("/transactions", controllers.CreateTransaction)
	r.Run()
}
