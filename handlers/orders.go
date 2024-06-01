package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/WallaceMartinz/ramenGO/data"
	"github.com/gin-gonic/gin"
)



// PostOrder handles the request to create a new order.
func PostOrder(ctx *gin.Context) {
	var reqBody data.OrderRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody); err != nil {
		handleGenericError(ctx)
		return
	}

	brothName, proteinName, err := validate(reqBody)
	if err != nil {
		HandleBadRequest(ctx)
		return
	}

	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleBadRequest(ctx)
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	apikey := ctx.GetHeader("x-api-key")
	if apikey == "" {
		HandleMissingAPIKey(ctx)
		return
	}

	req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", bytes.NewBuffer(bodyBytes))
	if err != nil {
		handleGenericError(ctx)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apikey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		handleGenericError(ctx)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		handleGenericError(ctx)
		return
	}

	response := createOrder(ctx, resp, brothName, proteinName)

	ctx.JSON(http.StatusCreated, response)
}

// createOrder processes the response from the external API and creates an OrderResponse.
func createOrder(ctx *gin.Context, resp *http.Response, brothName string, proteinName string) data.OrderResponse {
	var externalResponse data.ExternalResponse

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		handleGenericError(ctx)
	}

	if err := json.Unmarshal(respBody, &externalResponse); err != nil {
		handleGenericError(ctx)
	}

	imageURL := data.GetImageURL(proteinName)

	response := data.NewOrderResponse(
		externalResponse.OrderID,
		fmt.Sprintf("%s and %s Ramen", brothName, proteinName),
		imageURL,
	)

	return response
}

// Validate the required fields of the order request
func validate(reqBody data.OrderRequest) (string, string, error) {
	if reqBody.BrothID == "" || reqBody.ProteinID == "" {
		return "", "", errors.New("both brothId and proteinId are required")
	}
	brothName, err := GetBrothNameById(reqBody.BrothID)
	if err != nil {
		return "", "", err
	}

	proteinName, err := GetProteinNameById(reqBody.ProteinID)
	if err != nil {
		return "", "", err
	}
	return brothName, proteinName, nil
}