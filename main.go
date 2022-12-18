package main

import (
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/router"
)

func main() {
	config.InitConfig()
	router.Serve()
}
