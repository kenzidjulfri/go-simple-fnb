package response

type OrderResponse struct {
	OrderID  uint      `json:"orderId"`
	Printers []Printer `json:"printers"`
}
