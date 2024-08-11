package model

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	ID              uint    `gorm:"primaryKey"`
	Name            string  `gorm:"type:varchar(255)"`
	ProductID       uint    `gorm:"type:bigint"`
	Product         Product `gorm:"foreignKey:ProductID"`
	AdditionalPrice float64 `gorm:"additionalPrice"`
}
