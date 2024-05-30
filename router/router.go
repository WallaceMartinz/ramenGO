package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	router.Run() // Serve on 0.0.0.0:8080 
}
