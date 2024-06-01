package data

// OrderResponse represents the response payload for an order.
type OrderResponse struct {
	ID			string `json:"id"`
	Description string `json:"description"`
	Image		string `json:"image"`
}

// ExternalResponse represents the response from an external order ID generation service.
type ExternalResponse struct {
	OrderID 	string `json:"orderId"`
}

// getImageURL returns the image URL based on the protein name.
func GetImageURL(proteinName string) string {
	ramens := map[string]string{
		"Chasu":            "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
		"Yasai Vegetarian": "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png",
		"Karaague":         "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png",
	}

	return ramens[proteinName]
}

// NewOrderResponse creates a new OrderResponse instance.
func NewOrderResponse(orderID, description, image string) OrderResponse {
	return OrderResponse{
		ID:          orderID,
		Description: description,
		Image:       image,
	}
}