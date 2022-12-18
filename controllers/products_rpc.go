package controllers

import (
	"context"

	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/generated/products"
	"github.com/thirumathikart/thirumathikart-product-service/models"
)

type ProductRPCServer struct {
	products.UnimplementedProductServiceServer
}

func (ProductRPCServer) GetProductsRPC(ctx context.Context, request *products.GetProductsRequest) (*products.GetProductsResponse, error) {
	var err error
	db := config.GetDB()
	var response []*products.Product
	for _, id := range request.Products {
		var product models.Product
		db.First(&product, id)
		response = append(response, &products.Product{
			ProductId:    uint32(product.ID),
			SellerId:     uint32(product.SellerID),
			CategoryId:   uint32(product.CategoryID),
			ProductTitle: product.Title,
			ProductPrice: uint32(product.Price),
		})
	}
	return &products.GetProductsResponse{Products: response}, err
}
