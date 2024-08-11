package repository

import (
	"github.com/kenzidjulfri/model"
	"gorm.io/gorm"
)

// single responsibility principle
type ProductRepository interface {
	GetByID(db *gorm.DB, id uint) (model.Product, error)
	Create(db *gorm.DB, data *model.Product) error
	Update(db *gorm.DB, data *model.Product) error
	Delete(db *gorm.DB, id uint) error
	Restore(db *gorm.DB, id uint) error
}

var ProductRepo ProductRepository = nil

type ProductRepositoryImpl struct{}

func GetProductRepository() ProductRepository {
	if ProductRepo == nil {
		ProductRepo = &ProductRepositoryImpl{}
	}
	return ProductRepo
}

func (repo *ProductRepositoryImpl) GetByID(db *gorm.DB, id uint) (model.Product, error) {
	var data model.Product
	err := db.First(&data, id).Error
	return data, err
}

func (repo *ProductRepositoryImpl) Create(db *gorm.DB, data *model.Product) error {
	return db.Create(&data).Error
}

func (repo *ProductRepositoryImpl) Update(db *gorm.DB, data *model.Product) error {
	return db.Save(&data).Error
}

func (repo *ProductRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&model.Product{}, id).Error
}

func (repo *ProductRepositoryImpl) Restore(db *gorm.DB, id uint) error {
	return db.Update("deleted_at", nil).Error
}
