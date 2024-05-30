package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOrders(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST orders",
	})
}
