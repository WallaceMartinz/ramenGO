package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/WallaceMartinz/ramenGO/data"
	"github.com/gin-gonic/gin"
)



// GetBroths handles the request to get all broths.
func GetBroths(ctx *gin.Context) {
	broths, err := data.GetBroths()
	if err != nil {
		handleGenericError(ctx)
		return
	}
	ctx.JSON(http.StatusOK, broths)
}

// GetBrothNameById returns the name of the broth based on the provided ID.
func GetBrothNameById(brothId string) (string, error) {
	id, err := strconv.Atoi(brothId)
	if err != nil {
		return "", err
	}

	broths, err := data.GetBroths()
	if err != nil {
		return "", err
	}

	for _, broth := range broths {
		if broth.ID == id {
			return broth.Name, nil
		}
	}
	return "", errors.New("invalid broth ID")
}
