package routes

import (
	"smallbank/server/controllers"
	"smallbank/server/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	injectMiddlewares(router)
	initializeRoutes(router)
	return router
}

func injectMiddlewares(router *gin.Engine) {
	router.Use(middlewares.DBConnectionPool())
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/accounts/:id", controllers.GetAccount)
	router.GET("/accounts", controllers.GetAccountList)
	router.POST("/accounts", controllers.CreateAccount)
	router.PUT("/accounts/:id", controllers.UpdateAccount)
	router.DELETE("/accounts/:id", controllers.DeleteAccount)

	router.GET("/users/:id", controllers.GetUser)
	router.GET("/users", controllers.GetUserList)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.GET("/transfers/:UserId", controllers.GetTransferList)
	router.POST("/transfers", controllers.CreateTransfer)

	router.GET("/transactions/:UserId", controllers.GetTransactionList)
	router.POST("/transactions", controllers.CreateTransaction)
}
