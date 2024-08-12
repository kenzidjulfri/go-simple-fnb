package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"type:varchar(255)"`
	Category string    `gorm:"type:varchar(255)"`
	Price    float64   `gorm:"type:numeric(10,2)"`
	Variants []Variant `gorm:"foreignKey:ProductID"`
}
