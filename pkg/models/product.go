package models
import "gorm.io/gorm"
type Category struct {
	gorm.Model
	//ID           uint   `gorm:"primary key" json:"id"`
	CategoryName string `json:"category_name" validate:"required"`
	Description  string `json:"category_description" validate:"required"`
	ImageUrl     string `json:"category_imageUrl" validate:"required"`
}
type Product struct {
	gorm.Model
	//ID          uint     `gorm:"primary key" json:"id"`
	CategoryName           string     `json:"category_name" validate:"required"`
	//Category             Category `gorm:"foriegnkey:CategoryID;references:ID" json:"category,omitempty"`
	ProductName          string   `json:"product_name" validate:"required"`
	Description          string   `json:"product_description" validate:"required"`
	ImageUrl             string   `json:"product_imageUrl" validate:"required"`
	Price                float64  `gorm:"type:decimal(10,2)" json:"price" validate:"required"`
	Stock                uint     `json:"stock"`
	Popular              bool     `gorm:"type:boolean;default:false" json:"popular" validate:"required"`
	Size                 string   `gorm:"type:varchar(10); check:size IN ('Medium', 'Small', 'Large')" json:"size" validate:"required,oneof=Medium Small Large"`
	//HasOffer             bool     `gorm:"default:false"`
	//OfferDiscountPercent uint     `gorm:"default:0"`
}