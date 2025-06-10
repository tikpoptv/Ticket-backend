package router

import (
	"ticket-backend/internal/handlers/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Welcome message
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API!",
		})
	})

	// Register endpoint
	authHandler := auth.NewRegisterHandler()
	r.POST("/register", authHandler.Register)

	return r
}
