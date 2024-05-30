package schemas

type OrderRequest struct {
	BrothID   int `json:"brothId"`
	ProteinID int `json:"proteinId"`
}
