package controllers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

func ListProductsBySeller(sellerID int) []models.Product {
	db := config.GetDB()
	var products []models.Product
	db.Find(&products, "seller_id = ?", sellerID)
	return products
}

func ListProductsByCategory(categoryID int) []models.Product {
	db := config.GetDB()
	var products []models.Product
	db.Find(&products, "category_id = ?", categoryID)
	return products
}

func CreateProduct(c echo.Context) error {
	db := config.GetDB()
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	categoryID, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return err
	}
	sellerID, err := strconv.Atoi(c.FormValue("seller_id"))
	if err != nil {
		return err
	}
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return err
	}
	stock, err := strconv.Atoi(c.FormValue("stock"))
	if err != nil {
		return err
	}
	product := models.Product{
		Title:       c.FormValue("title"),
		CategoryID:  categoryID,
		SellerID:    sellerID,
		Price:       price,
		Description: c.FormValue("description"),
		Stock:       stock,
	}
	db.Create(&product)
	files := form.File["files"]
	err = UploadProductImage(files, product.ID, db)
	if err != nil {
		return err
	}
	return nil
}
