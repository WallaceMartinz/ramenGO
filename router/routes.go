package router

import (
	"github.com/WallaceMartinz/ramenGO/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api")
	{
		v1.GET("/broths", handler.GetBroths)

		v1.GET("/proteins", handler.GetProteins)

		v1.POST("/orders", handler.PostOrders)
	}
}
