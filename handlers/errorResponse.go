package handlers

import (
	"net/http"

	"github.com/WallaceMartinz/ramenGO/data"
	"github.com/gin-gonic/gin"
)

// handleGenericError sends a generic error response with status 500.
func handleGenericError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, data.ErrorResponse{
		Error: "could not place order",
	})
}

// handleBadRequest sends a bad request error response with status 400.
func HandleBadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, data.ErrorResponse{
		Error: "both brothId and proteinId are required",
	})
}

// HandleMissingAPIKey sends a forbidden error response with status 403.
func HandleMissingAPIKey(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, data.ErrorResponse{
		Error: "x-api-key header missing",
	})
}
