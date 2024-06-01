package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/WallaceMartinz/ramenGO/data"
	"github.com/gin-gonic/gin"
)

// GetProteins handles the request to get all proteins.
func GetProteins(ctx *gin.Context) {
	proteins, err := data.GetProteins()

	if err != nil {
		handleGenericError(ctx)
		return
	}
	ctx.JSON(http.StatusOK, proteins)
}

// GetProteinNameById returns the name of the protein based on the provided ID.
func GetProteinNameById(proteinId string) (string, error) {
	id, err := strconv.Atoi(proteinId)
	if err != nil {
		return "", err
	}

	proteins, err := data.GetProteins()
	if err != nil {
		return "", err
	}

	for _, protein := range proteins {
		if protein.ID == id {
			return protein.Name, nil
		}
	}
	return "", errors.New("invalid protein ID")
}
