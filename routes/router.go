package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/controllers"
)

func Init(e *echo.Echo) {
	// Static files
	e.Static("/static", "product_images")

	// Routes
	e.POST("/create_product", controllers.CreateProduct)
}
