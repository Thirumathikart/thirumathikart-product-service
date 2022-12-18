package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

func ListProductsBySeller(c echo.Context) error {
	db := config.GetDB()
	sellerID := c.QueryParam("seller")
	if sellerID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	var products []models.ProductImage
	db.Preload("Product","seller_id = ?",sellerID).Find(&products)
	return c.JSONPretty(http.StatusOK, products, "  ")
}

func ListProductsByCategory(c echo.Context) error {
	db := config.GetDB()
	categoryID := c.QueryParam("category")
	if categoryID == "" {
		return echo.NewHTTPError(400, "Category ID is required")
	}
	var products []models.ProductImage
	db.Preload("Product","category_id = ?",categoryID).Find(&products)
	log.Printf("%d rows found.", len(products))
	fmt.Println(products)
	return c.JSONPretty(http.StatusOK, products," ")
}
