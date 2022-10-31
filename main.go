package main

import (
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/models"
	"github.com/thirumathikart/thirumathikart-product-service/routes"
	"github.com/thirumathikart/thirumathikart-product-service/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitConfig()

	config.ConnectDB()
	models.MigrateDB()

	server := echo.New()
	utils.InitLogger(server)
	server.Use(middleware.CORS())

	routes.Init(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
