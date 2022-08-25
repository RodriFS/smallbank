package main

import (
	"smallbank/main/constants"
	"smallbank/main/controllers"
	"smallbank/main/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	constants.LoadConfig()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	r.POST("/account", controllers.CreateAccount)

	r.POST("/user", controllers.CreateUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
