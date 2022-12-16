package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

type productsResponse struct {
	products []models.Product
}

type productResponse struct {
	product models.Product
}

func ListProductsBySeller(c echo.Context) error {
	db := config.GetDB()
	sellerID := c.QueryParam("seller")
	if sellerID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	var products []models.Product
	db.Find(&products, "seller_id = ?", sellerID)
	productsList := productsResponse{
		products: products,
	}
	return c.JSONPretty(http.StatusOK, productsList, "  ")
}

func ListProductsByCategory(c echo.Context) error {
	db := config.GetDB()
	categoryID := c.QueryParam("category")
	if categoryID == "" {
		return echo.NewHTTPError(400, "Category ID is required")
	}
	var products []models.Product
	db.Find(&products, "category_id = ?", categoryID)
	productsList := productsResponse{
		products: products,
	}
	return c.JSONPretty(http.StatusOK, productsList, "  ")
}

func UpdateProductStock(c echo.Context) error {
	db := config.GetDB()
	productID := c.Param("id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	stock, err := strconv.Atoi(c.Param("stock"))
	if err != nil {
		return err
	}
	var product models.Product
	db.First(&product, productID)
	product.Stock = stock
	db.Save(&product)
	return c.HTML(http.StatusOK, "")
}

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
	return c.HTML(http.StatusOK, "")
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
	productDetails := productResponse{
		product: product,
	}
	return c.JSONPretty(http.StatusOK, productDetails, "  ")
}

func UpdateProductPrice(c echo.Context) error {
	db := config.GetDB()
	// Check if User is seller
	productID := c.Param("product_id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	price, err := strconv.Atoi(c.Param("price"))
	if err != nil {
		return err
	}
	var product models.Product
	db.First(&product, productID)
	product.Price = price
	db.Save(&product)
	return c.HTML(http.StatusOK, "")
}
