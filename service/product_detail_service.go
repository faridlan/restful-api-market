package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ProductDetailService interface {
	FindProduct(ctx context.Context, productId int) web.ProductResponse
	AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse
}
