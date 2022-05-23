package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindyId(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	CreateImg(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
}
