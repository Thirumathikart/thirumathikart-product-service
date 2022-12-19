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
	sellerID, err := strconv.Atoi(c.QueryParam("seller"))
	if err != nil {
		return err
	}
	var products []models.ProductImage
	if err:=db.Joins("JOIN products ON products.id = product_images.product_id").Where("products.seller_id = ?",sellerID).Preload("Product").Find(&products).Error;err!=nil{
		return c.JSON(http.StatusInternalServerError," ")
	}
	return c.JSON(http.StatusOK, products)
}

func ListProductsByCategory(c echo.Context) error {
	db := config.GetDB()
	categoryID, err := strconv.Atoi(c.QueryParam("category"))
	if err != nil {
		return err
	}
	var products []models.ProductImage
	if err:=db.Joins("JOIN products ON products.id = product_images.product_id").Where("products.category_id = ?",categoryID).Preload("Product").Find(&products).Error;err!=nil{
		return c.JSON(http.StatusInternalServerError," ")
	}
	//db.Preload("Product","category_id = ?",categoryID).Find(&products)
	log.Printf("%d rows found.", len(products))
	fmt.Println(products)
	return c.JSONPretty(http.StatusOK, products," ")
}
