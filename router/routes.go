package router

import (
	"github.com/WallaceMartinz/ramenGO/handlers"
	"github.com/gin-gonic/gin"
)

// Initializes routes for the application.
func initializeRoutes(router *gin.Engine) {
	r := router.Group("/")
	{
		r.GET("/broths", handlers.GetBroths)

		r.GET("/proteins", handlers.GetProteins)

		r.POST("/order", handlers.PostOrder)
	}
}
