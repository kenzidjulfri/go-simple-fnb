package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kenzidjulfri/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = database
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.Printer{},
		&model.Table{},
		&model.Product{},
		&model.Variant{},
		&model.PromoDetail{},
		&model.Promo{},
		&model.OrderDetail{},
		&model.Order{},
	); err != nil {
		return err
	}

	return nil
}
