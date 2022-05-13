package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindAll(ctx context.Context) []web.ProductResponse
}
