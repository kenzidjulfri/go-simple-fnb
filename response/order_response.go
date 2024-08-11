package response

type OrderResponse struct {
	Printers []Printer `json:"printers"`
}

type Printer struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
