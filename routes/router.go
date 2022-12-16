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
	e.GET("/list_products_by_seller", controllers.ListProductsBySeller)
	e.GET("/list_products_by_category", controllers.ListProductsByCategory)
	e.POST("/update_product_stock/", controllers.UpdateProductStock)
	e.POST("/delete_product", controllers.DeleteProduct)
	e.POST("/update_product_price", controllers.UpdateProductPrice)
	e.POST("/get_product_details", controllers.GetProductDetails)
}
