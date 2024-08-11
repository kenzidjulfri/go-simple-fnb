package response

type BillResponse struct {
	Table      string               `json:"table"`
	Details    []BillDetailResponse `json:"details"`
	TotalPrice float64              `json:"totalPrice"`
}

type BillDetailResponse struct {
	ItemName  string  `json:"itemName"`
	Qty       uint    `json:"qty"`
	BasePrice float64 `json:"basePrice"`
	SumPrice  float64 `json:"sumPrice"`
}
