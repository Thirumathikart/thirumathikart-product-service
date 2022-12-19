package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

func ListProductsBySeller(c echo.Context) error {
	db := config.GetDB()
	sellerIDParam := c.QueryParam("seller")
	if sellerIDParam == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	sellerID, err := strconv.Atoi(sellerIDParam)
	if err != nil {
		log.Panicln(err)
	}

	var products []models.ProductImage
	db.Preload("Product", "seller_id = ?", sellerID).Find(&products)
	return c.JSONPretty(http.StatusOK, products, "  ")
}

func ListProductsByCategory(c echo.Context) error {
	db := config.GetDB()
	categoryIDParam := c.QueryParam("category")
	log.Println("categroryID", categoryIDParam)
	if categoryIDParam == "" {
		return echo.NewHTTPError(400, "Category ID is required")
	}
	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		log.Panicln(err)
	}

	var products []models.ProductImage
	db.Where("Product", "category_id = ?", categoryID).Find(&products)
	log.Printf("%d rows found.", len(products))
	fmt.Println(products)
	return c.JSONPretty(http.StatusOK, products, " ")
}
