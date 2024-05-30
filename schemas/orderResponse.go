package schemas

type OrderResponse struct {
	ID			int    `json:"id"`
	Description string `json:"description"`
	Image		string `json:"image"`
}