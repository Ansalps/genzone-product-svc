package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Ansalps/genzone-product-svc/pkg/db"
	"github.com/Ansalps/genzone-product-svc/pkg/models"
	"github.com/Ansalps/genzone-product-svc/pkg/pb"
	"gorm.io/gorm"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	var category models.Category
	fmt.Println("req", req)
	category.CategoryName = req.Categoryname
	category.Description = req.Description
	category.ImageUrl = req.Imageurl
	if result := s.H.DB.Create(&category); result.Error != nil {
		return &pb.CreateCategoryResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateCategoryResponse{
		Status: http.StatusCreated,
		Id:     int64(category.ID),
	}, nil
}
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product
	fmt.Println("req", req)
	product.CategoryName = req.Categoryname
	product.ProductName = req.Productname
	product.Description = req.Description
	product.ImageUrl = req.Imageurl
	product.Price = req.Price
	product.Stock = uint(req.Stock)
	product.Popular = req.Popular
	product.Size = req.Size
	var category models.Category
	if err := s.H.DB.Where("category_name = ?", product.CategoryName).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.CreateProductResponse{
				Status: http.StatusBadRequest,
				Error:  "Such category does not exist",
			}, nil
		}
	}
	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     int64(product.ID),
	}, nil
}
func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	productID := req.GetProductId()
	log.Printf("Received request for product ID: %s", productID)
	var product models.Product
	if err := s.H.DB.Where("id = ?", product.ID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.GetProductResponse{
				Status: http.StatusBadRequest,
				Error:  "Such product id does not exist",
			}, nil
		}
	}
	return &pb.GetProductResponse{
		Productid: productID,
		Categoryname: product.CategoryName,
		Productname: product.ProductName,
		Description: product.Description,
		Imageurl: product.ImageUrl,
		Price: product.Price,
		Stock: int64(product.Stock),
		Popular: product.Popular,
	}, nil

}
