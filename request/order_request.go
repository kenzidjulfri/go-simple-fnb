package request

type OrderRequest struct {
	TableID uint                 `json:"tableId" binding:"required"`
	Details []OrderDetailRequest `json:"details" binding:"required"`
}

type OrderDetailRequest struct {
	ItemID   uint   `json:"itemId" binding:"required"`
	ItemType string `json:"itemType" binding:"required"`
	Qty      uint   `json:"qty" binding:"required"`
}
