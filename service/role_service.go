package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type RoleService interface {
	Create(ctx context.Context, request web.RoleCreateRequest) web.RoleResponse
	Update(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse
	FindAll(ctx context.Context) []web.RoleResponse
}
