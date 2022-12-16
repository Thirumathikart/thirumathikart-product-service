package controllers

import (
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

type ListProductsResponse struct {
	products []models.Product
}

type SingleProductResponse struct {
	product models.Product
}
