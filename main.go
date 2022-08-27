package main

import (
	"smallbank/server/initializers"
	"smallbank/server/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	gin := routes.SetupRouter()
	gin.Run()
}
