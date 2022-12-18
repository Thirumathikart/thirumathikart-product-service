package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

func CreateProduct(c echo.Context) error {
	db := config.GetDB()
	//  Check if User is seller
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
		log.Println(err)
		return err
	}
	stock, err := strconv.Atoi(c.FormValue("stock"))
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

func DeleteProduct(c echo.Context) error {
	db := config.GetDB()
	// Check if User is seller
	productID := c.Param("product_id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	var product models.Product
	db.First(&product, productID)
	db.Delete(&product)
	return c.HTML(http.StatusOK, "")
}

func GetProductDetails(c echo.Context) error {
	db := config.GetDB()
	productID := c.FormValue("product_id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	var product models.Product
	db.First(&product, productID)
	productDetails := SingleProductResponse{
		product: product,
	}
	return c.JSONPretty(http.StatusOK, productDetails, "  ")
}
