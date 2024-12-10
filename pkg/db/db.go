package db

import (
	"log"

	"github.com/Ansalps/genzone-product-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	return Handler{db}
}
