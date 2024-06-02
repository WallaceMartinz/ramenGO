package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WallaceMartinz/ramenGO/data"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProteinNameById(t *testing.T) {
	tests := []struct {
		ProteinId    string
		expectedName string
		expectedErr  error
	}{
		{"1", "Chasu", nil},
		{"2", "Yasai Vegetarian", nil},
		{"3", "Karaague", nil},
		{"4", "", errors.New("invalid protein ID")},
		{"invalid", "", errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax")},
	}

	for _, tt := range tests {
		name, err := GetProteinNameById(tt.ProteinId)
		if tt.expectedErr != nil {
			assert.EqualError(t, err, tt.expectedErr.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tt.expectedName, name)
	}
}

func TestGetBrothNameById(t *testing.T) {
	tests := []struct {
		brothId      string
		expectedName string
		expectedErr  error
	}{
		{"1", "Salt", nil},
		{"2", "Shoyu", nil},
		{"3", "Miso", nil},
		{"4", "", errors.New("invalid broth ID")},
		{"invalid", "", errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax")},
	}

	for _, tt := range tests {
		name, err := GetBrothNameById(tt.brothId)
		if tt.expectedErr != nil {
			assert.EqualError(t, err, tt.expectedErr.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tt.expectedName, name)
	}
}

func TestPostOrder_HandleMissingAPIKey(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	mockRequestBody := bytes.NewBufferString(`{"brothId": "1", "proteinId": "1"}`)
	ctx.Request = httptest.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", mockRequestBody)
	ctx.Request.Header.Set("Content-Type", "application/json")

	PostOrder(ctx)

	assert.Equal(t, http.StatusForbidden, ctx.Writer.Status())
}

func TestPostOrder_HandleGenerateIDRequestFailure(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	mockRequestBody := bytes.NewBufferString(`{"brothId": "1", "proteinId": "1"}`)
	ctx.Request = httptest.NewRequest("POST", "/", mockRequestBody)
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("x-api-key", "valid-key")

	PostOrder(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
}

func TestPostOrder_HandleBadRequestFailure(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	mockRequestBody := bytes.NewBufferString(`{"brothId": "", "proteinId": ""}`)
	ctx.Request = httptest.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", mockRequestBody)
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("x-api-key", "valid-key")

	PostOrder(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestCreateOrder_DecodeResponseData(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	externalResponseData := `{"OrderID": "12345"}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(externalResponseData)),
	}

	response := createOrder(ctx, resp, "Salt", "Chasu")

	assert.Equal(t, "12345", response.ID)
	assert.Equal(t, "Salt and Chasu Ramen", response.Description)
	assert.Equal(t, "https://tech.redventures.com.br/icons/ramen/ramenChasu.png", response.Image)
}

func TestCreateOrder_DecodeResponseDataError(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("invalid json")),
	}

	response := createOrder(ctx, resp, "Salt", "Chasu")

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	assert.Empty(t, response.ID)
	assert.Empty(t, response.Description)
	assert.Empty(t, response.Image)
}

func TestCreateOrder_GetImageURLError(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewBufferString(`{"OrderID": "12345"}`)),
	}

	response := createOrder(ctx, resp, "Salt", "InvalidProtein")

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	assert.Empty(t, response.ID)
	assert.Empty(t, response.Description)
	assert.Empty(t, response.Image)
}

func TestValidate(t *testing.T) {
	// Mock de erro ao não fornecer BrothID e ProteinID
	mockError := "both brothId and proteinId are required"
	brothName, proteinName, err := validate(data.OrderRequest{})
	assert.Equal(t, "", brothName)
	assert.Equal(t, "", proteinName)
	assert.EqualError(t, err, mockError)

	// Entradas corretas
	brothName, proteinName, err = validate(data.OrderRequest{
		BrothID:   "1",
		ProteinID: "2",
	})
	assert.Equal(t, "Salt", brothName)
	assert.Equal(t, "Yasai Vegetarian", proteinName)
	assert.NoError(t, err)
}

