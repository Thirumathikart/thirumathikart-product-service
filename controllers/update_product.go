package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/middlewares"
	"github.com/thirumathikart/thirumathikart-product-service/models"
	"github.com/thirumathikart/thirumathikart-product-service/utils"
)

func UpdateProduct(c echo.Context) error {
	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.SendResponse(c, http.StatusBadRequest, "Bad Request")
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	categoryID, err := strconv.Atoi(c.FormValue("category_id"))
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
	productID, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Println(err)
		return err
	}
	db := config.GetDB()
	var product models.Product
	res := db.Where("id = ?", productID).First(&product)
	if res.Error != nil {
		log.Println(res.Error)
		return c.JSON(http.StatusBadGateway, "Bad Request")
	}
	log.Println("product:", product)
	log.Println("userDetails:", userDetails)
	if product.SellerID != int(userDetails.UserId) {
		log.Println(res.Error)
		return c.JSON(http.StatusUnauthorized, "Seller Unauthorized")
	}
	res = db.Model(&product).Updates(
		models.Product{
			Title:       c.FormValue("title"),
			CategoryID:  categoryID,
			Price:       price,
			Description: c.FormValue("description"),
			Stock:       stock,
		})
	if res.Error != nil {
		log.Println(res.Error)
		return c.JSON(http.StatusBadGateway, "Bad Request")
	}
	db.Create(&product)
	files := form.File["files"]
	err = UpdateProductImages(files, product.ID, db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
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
