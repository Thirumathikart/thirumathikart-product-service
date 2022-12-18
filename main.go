package main

import (
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/routes"
)

func main() {
	config.InitConfig()
	routes.Serve()
}
