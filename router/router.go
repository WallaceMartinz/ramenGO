package router

import (
	"github.com/WallaceMartinz/ramenGO/middleware"
	"github.com/gin-gonic/gin"
)

// Init initializes the Gin router, sets up middlewares, defines routes, and starts the server.
func Init() {
	router := gin.Default()

	router.Use(
		gin.HandlerFunc(middleware.CorsMiddleware()),
		gin.HandlerFunc(middleware.ApiKeyMiddleware()),
	)

	initializeRoutes(router)

	router.Run()
}
