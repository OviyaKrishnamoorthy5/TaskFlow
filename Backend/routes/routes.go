package routes

import (
	"taskflow-backend/controllers"
	"taskflow-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// public routes
	r.GET("/ping", controllers.Ping)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 🔐 protected test route
	r.GET("/profile", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are inside protected route 🔐",
		})
	})

	return r
}
