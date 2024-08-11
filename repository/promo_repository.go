package repository

import (
	"github.com/kenzidjulfri/model"
	"gorm.io/gorm"
)

// single responsibility principle
type PromoRepository interface {
	GetByID(db *gorm.DB, id uint) (model.Promo, error)
	Create(db *gorm.DB, data *model.Promo) error
	Update(db *gorm.DB, data *model.Promo) error
	Delete(db *gorm.DB, id uint) error
	Restore(db *gorm.DB, id uint) error
}

var PromoRepo PromoRepository = nil

type PromoRepositoryImpl struct{}

func GetPromoRepository() PromoRepository {
	if PromoRepo == nil {
		PromoRepo = &PromoRepositoryImpl{}
	}
	return PromoRepo
}

func (repo *PromoRepositoryImpl) GetByID(db *gorm.DB, id uint) (model.Promo, error) {
	var data model.Promo
	err := db.Preload("PromoDetails").Preload("PromoDetails.Variant").Preload("PromoDetails.Product").First(&data, id).Error
	return data, err
}

func (repo *PromoRepositoryImpl) Create(db *gorm.DB, data *model.Promo) error {
	return db.Create(&data).Error
}

func (repo *PromoRepositoryImpl) Update(db *gorm.DB, data *model.Promo) error {
	return db.Save(&data).Error
}

func (repo *PromoRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&model.Promo{}, id).Error
}

func (repo *PromoRepositoryImpl) Restore(db *gorm.DB, id uint) error {
	return db.Update("deleted_at", nil).Error
}
