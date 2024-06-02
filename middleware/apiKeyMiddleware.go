package middleware

import (
	"github.com/WallaceMartinz/ramenGO/handlers"
	"github.com/gin-gonic/gin"
)

// ApiKeyMiddleware is a function to validate the presence of an API key in the request headers.
func ApiKeyMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        apikey := ctx.GetHeader("x-api-key")
        if apikey == "" {
            handlers.HandleMissingAPIKey(ctx)
            return
        }
        ctx.Next()
    }
}
