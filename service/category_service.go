package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
