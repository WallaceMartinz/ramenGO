package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBroths(t *testing.T) {
	expectedBroths := []Broth{
		{
			ID:            1,
			ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
			Name:          "Salt",
			Description:   "Simple like the seawater, nothing more",
			Price:         10,
		},
		{
			ID:            2,
			ImageInactive: "https://tech.redventures.com.br/icons/shoyu/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/shoyu/active.svg",
			Name:          "Shoyu",
			Description:   "The good old and traditional soy sauce",
			Price:         10,
		},
		{
			ID:            3,
			ImageInactive: "https://tech.redventures.com.br/icons/miso/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/miso/active.svg",
			Name:          "Miso",
			Description:   "Paste made of fermented soybeans",
			Price:         12,
		},
	}

	broths, err := GetBroths()
	assert.NoError(t, err)
	assert.Equal(t, expectedBroths, broths)
}

func TestGetProtein(t *testing.T) {
	expectedProteins := []Protein{
		{
			ID:            1,
			ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
			Name:          "Chasu",
			Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
			Price:         10,
		},
		{
			ID:            2,
			ImageInactive: "https://tech.redventures.com.br/icons/yasai/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/yasai/active.svg",
			Name:          "Yasai Vegetarian",
			Description:   "A delicious vegetarian lamen with a selection of season vegetables.",
			Price:         10,
		},
		{
			ID:            3,
			ImageInactive: "https://tech.redventures.com.br/icons/chicken/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/chicken/active.svg",
			Name:          "Karaague",
			Description:   "Three units of fried chicken, moyashi, ajitama egg and other vegetables.",
			Price:         12,
		},
	}

	proteins, err := GetProteins()
	assert.NoError(t, err)
	assert.Equal(t, expectedProteins, proteins)
}

func TestGetImageURL(t *testing.T) {
	tests := []struct {
		proteinName string
		expectedURL string
	}{
		{"Chasu", "https://tech.redventures.com.br/icons/ramen/ramenChasu.png"},
		{"Yasai Vegetarian", "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png"},
		{"Karaague", "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png"},
		{"Unknown", ""},
	}

	for _, tt := range tests {
		url := GetImageURL(tt.proteinName)
		assert.Equal(t, tt.expectedURL, url)
	}
}

func TestNewOrderResponse(t *testing.T) {
	orderID := "12345"
	description := "Salt and Chasu Ramen"
	image := "https://tech.redventures.com.br/icons/ramen/ramenChasu.png"
	expectedResponse := OrderResponse{
		ID:          orderID,
		Description: description,
		Image:       image,
	}

	response := NewOrderResponse(orderID, description, image)
	assert.Equal(t, expectedResponse, response)
}


