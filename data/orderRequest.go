package data

// OrderRequest represents the request payload for creating a new order.
type OrderRequest struct {
	BrothID   string `json:"brothId"`
	ProteinID string `json:"proteinId"`
}
