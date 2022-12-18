package routes

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/controllers"
	"github.com/thirumathikart/thirumathikart-product-service/generated/products"
	"github.com/thirumathikart/thirumathikart-product-service/middlewares"
	"google.golang.org/grpc"
)

func Serve() {
	// Static files
	var wg sync.WaitGroup
	wg.Add(2)

	// REST server
	httpPort := config.ServerPort
	e := echo.New()
	middlewares.InitLogger(e)
	e.Use(middleware.CORS())
	e.Static("/static", "product_images")

	// Routes
	e.POST("/create_product", controllers.CreateProduct)
	e.GET("/list_products_by_seller", controllers.ListProductsBySeller)
	e.GET("/list_products_by_category", controllers.ListProductsByCategory)
	e.POST("/update_product_stock/", controllers.UpdateProductStock)
	e.POST("/delete_product", controllers.DeleteProduct)
	e.POST("/update_product_price", controllers.UpdateProductPrice)
	e.POST("/get_product_details", controllers.GetProductDetails)
	e.POST("/update_product_title", controllers.UpdateProductTitle)
	e.POST("update_product_description", controllers.UpdateProductDescription)
	go func() {
		e.Logger.Fatal(e.Start(":" + httpPort))
	}()

	// GRPC server
	grpcPort := config.RPCPort
	grpcServer := grpc.NewServer(middlewares.WithServerUnaryInterceptor())
	products.RegisterProductServiceServer(grpcServer, &controllers.ProductRPCServer{})

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			log.Panic("grpc server running error on", err)
		}
		err1 := grpcServer.Serve(lis)
		if err1 != nil {
			log.Panic("grpc server running error on", err1)
		}
	}()

	wg.Wait()
}
