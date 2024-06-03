package middleware

import (
	"os"

	"github.com/WallaceMartinz/ramenGO/handlers"
	"github.com/gin-gonic/gin"
)

// ApiKeyMiddleware is a function to validate the presence of an API key in the request headers.
func ApiKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikeyFromHeader := ctx.GetHeader("x-api-key")
		if apikeyFromHeader == "" {
			handlers.HandleMissingAPIKey(ctx)
			return
		}

		apikeyFromEnv := os.Getenv("X_API_KEY")
		if apikeyFromEnv == "" {
			handlers.HandleMissingAPIKey(ctx)
			return
		}

		if apikeyFromEnv != apikeyFromHeader {
			handlers.HandleMissingAPIKey(ctx)
			return
		}

		ctx.Next()
	}
}
