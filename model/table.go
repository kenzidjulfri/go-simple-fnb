package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}
