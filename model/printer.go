package model

import "gorm.io/gorm"

type Printer struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"type:varchar(10)"`
	Name  string `gorm:"type:varchar(255)"`
	Usage string `gorm:"type:varchar(255)"`
}
