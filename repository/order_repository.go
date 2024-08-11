package repository

import (
	"github.com/kenzidjulfri/model"
	"gorm.io/gorm"
)

// single responsibility principle
type OrderRepository interface {
	GetByID(db gorm.DB, id uint) (model.Order, error)
	Create(db gorm.DB, data *model.Order) error
	Update(db gorm.DB, data *model.Order) error
	Delete(db gorm.DB, id uint) error
	Restore(db gorm.DB, id uint) error
}

var OrderRepo OrderRepository

type OrderRepositoryImpl struct{}

func GetOrderRepository() OrderRepository {
	if OrderRepo == nil {
		OrderRepo = &OrderRepositoryImpl{}
	}
	return OrderRepo
}

func (repo *OrderRepositoryImpl) GetByID(db gorm.DB, id uint) (model.Order, error) {
	var data model.Order
	err := db.Preload("Table").First(&data, id).Error
	return data, err
}

func (repo *OrderRepositoryImpl) Create(db gorm.DB, data *model.Order) error {
	return db.Create(&data).Error
}

func (repo *OrderRepositoryImpl) Update(db gorm.DB, data *model.Order) error {
	return db.Save(&data).Error
}

func (repo *OrderRepositoryImpl) Delete(db gorm.DB, id uint) error {
	return db.Delete(&model.Order{}, id).Error
}

func (repo *OrderRepositoryImpl) Restore(db gorm.DB, id uint) error {
	return db.Update("deleted_at", nil).Error
}
