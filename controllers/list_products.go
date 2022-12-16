package controllers

import (
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
	var products []models.Product
	db.Find(&products, "seller_id = ?", sellerID)
	productsList := ListProductsResponse{
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
	productsList := ListProductsResponse{
		products: products,
	}
	return c.JSONPretty(http.StatusOK, productsList, "  ")
}
