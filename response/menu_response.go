package response

type MenuResponse struct {
	Category string            `json:"category"`
	Items    []SubMenuResponse `json:"items"`
}

type SubMenuResponse struct {
	ItemID    uint    `json:"itemId"`
	ItemName  string  `json:"itemName"`
	ItemType  string  `json:"itemType"`
	BasePrice float64 `json:"basePrice"`
}
