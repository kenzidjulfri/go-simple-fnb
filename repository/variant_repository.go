package repository

import (
	"github.com/kenzidjulfri/model"
	"gorm.io/gorm"
)

// single responsibility principle
type VariantRepository interface {
	GetByID(db gorm.DB, id uint) (model.Variant, error)
	Create(db gorm.DB, data *model.Variant) error
	Update(db gorm.DB, data *model.Variant) error
	Delete(db gorm.DB, id uint) error
	Restore(db gorm.DB, id uint) error
}

var VariantRepo VariantRepository

type VariantRepositoryImpl struct{}

func GetVariantRepository() VariantRepository {
	if VariantRepo == nil {
		VariantRepo = &VariantRepositoryImpl{}
	}
	return VariantRepo
}

func (repo *VariantRepositoryImpl) GetByID(db gorm.DB, id uint) (model.Variant, error) {
	var data model.Variant
	err := db.Preload("Product").First(&data, id).Error
	return data, err
}

func (repo *VariantRepositoryImpl) Create(db gorm.DB, data *model.Variant) error {
	return db.Create(&data).Error
}

func (repo *VariantRepositoryImpl) Update(db gorm.DB, data *model.Variant) error {
	return db.Save(&data).Error
}

func (repo *VariantRepositoryImpl) Delete(db gorm.DB, id uint) error {
	return db.Delete(&model.Variant{}, id).Error
}

func (repo *VariantRepositoryImpl) Restore(db gorm.DB, id uint) error {
	return db.Update("deleted_at", nil).Error
}
