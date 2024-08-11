package model

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	ID      uint          `gorm:"primaryKey"`
	Name    string        `gorm:"type:varchar(255)"`
	Details []PromoDetail `gorm:"foreignKey:PromoID"`
	Price   float64       `gorm:"type:numeric(10,2)"`
}

type PromoDetail struct {
	gorm.Model
	ID        uint     `gorm:"primaryKey"`
	PromoID   uint     `gorm:"type:bigint"`
	Promo     Promo    `gorm:"foreignKey:PromoID"`
	VariantID *uint    `gorm:"type:bigint"`
	Variant   *Variant `gorm:"foreignKey:VariantID"`
	ProductID uint     `gorm:"type:bigint"`
	Product   Product  `gorm:"foreignKey:ProductID"`
}
