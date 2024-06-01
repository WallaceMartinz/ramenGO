package middleware

import "github.com/gin-gonic/gin"

// CorsMiddleware is a function to handle Cross-Origin (CORS) headers.
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "https://tech.redventures.com.br")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, x-api-key")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}

		ctx.Next()
	}
}
