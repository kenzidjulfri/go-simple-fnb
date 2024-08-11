package repository

import (
	"github.com/kenzidjulfri/model"
	"gorm.io/gorm"
)

// single responsibility principle
type PrinterRepository interface {
	GetByID(db gorm.DB, id uint) (model.Printer, error)
	GetByUsage(db gorm.DB, usage string) (model.Printer, error)
	Create(db gorm.DB, data *model.Printer) error
	Update(db gorm.DB, data *model.Printer) error
	Delete(db gorm.DB, id uint) error
	Restore(db gorm.DB, id uint) error
}

var PrinterRepo PrinterRepository

type PrinterRepositoryImpl struct{}

func GetPrinterRepository() PrinterRepository {
	if PrinterRepo == nil {
		PrinterRepo = &PrinterRepositoryImpl{}
	}
	return PrinterRepo
}

func (repo *PrinterRepositoryImpl) GetByID(db gorm.DB, id uint) (model.Printer, error) {
	var data model.Printer
	err := db.First(&data, id).Error
	return data, err
}

func (repo *PrinterRepositoryImpl) GetByUsage(db gorm.DB, usage string) (model.Printer, error) {
	var data model.Printer
	err := db.Where("usage = ?", usage).First(&data).Error
	return data, err
}

func (repo *PrinterRepositoryImpl) Create(db gorm.DB, data *model.Printer) error {
	return db.Create(&data).Error
}

func (repo *PrinterRepositoryImpl) Update(db gorm.DB, data *model.Printer) error {
	return db.Save(&data).Error
}

func (repo *PrinterRepositoryImpl) Delete(db gorm.DB, id uint) error {
	return db.Delete(&model.Printer{}, id).Error
}

func (repo *PrinterRepositoryImpl) Restore(db gorm.DB, id uint) error {
	return db.Update("deleted_at", nil).Error
}
