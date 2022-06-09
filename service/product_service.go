package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpateRequest) web.ProductResponse
	Delete(ctx context.Context, productId string)
	FindyId(ctx context.Context, productId string) web.ProductResponse
	FindAll(ctx context.Context, pagination domain.Pagination) []web.ProductResponse
	CreateImg(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	FindSeeder(ctx context.Context, pagination domain.Pagination) []web.ProductResponse
}
