package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

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

func UpdateProductTitle(c echo.Context) error {
	db := config.GetDB()
	// Check if User is seller
	productID := c.Param("product_id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	title := c.Param("title")
	var product models.Product
	db.First(&product, productID)
	product.Title = title
	db.Save(&product)
	return c.HTML(http.StatusOK, "")
}

func UpdateProductDescription(c echo.Context) error {
	db := config.GetDB()
	// Check if User is seller
	productID := c.Param("product_id")
	if productID == "" {
		return echo.NewHTTPError(400, "Product ID is required")
	}
	description := c.Param("description")
	var product models.Product
	db.First(&product, productID)
	product.Description = description
	db.Save(&product)
	return c.HTML(http.StatusOK, "")
}
