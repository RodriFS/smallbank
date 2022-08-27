package middlewares

import (
	"smallbank/server/db"

	"github.com/gin-gonic/gin"
)

func DBConnectionPool() gin.HandlerFunc {
	pool := db.Init()
	return func(c *gin.Context) {
		c.Set("DB", pool)
		c.Next()
	}
}
