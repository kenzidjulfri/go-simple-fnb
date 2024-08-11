package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID       uint          `gorm:"primaryKey"`
	TableID  uint          `gorm:"type:bigint"`
	Table    Table         `gorm:"foreignKey:TableID"`
	Details  []OrderDetail `gorm:"foreignKey:OrderID"`
	IsClosed bool          `gorm:"type:bool;default:false"`
}

type OrderDetail struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"bigint"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ItemID    uint    `gorm:"type:bigint"`
	ItemName  string  `gorm:"type:varchar(255)"`
	ItemType  string  `gorm:"type:varchar(255)"`
	BasePrice float64 `gorm:"type:numeric(10,2)"`
	Qty       uint    `gorm:"type:smallint"`
	SumPrice  float64 `gorm:"type:numeric(10,2)"`
}
