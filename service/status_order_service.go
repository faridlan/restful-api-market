package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type StatusOrderService interface {
	Create(ctx context.Context, request web.StatusOrderCreate) web.StatusOrderResponse
	Update(ctx context.Context, request web.StatusOrderUpdate) web.StatusOrderResponse
	Delete(ctx context.Context, statusId string)
	FindById(ctx context.Context, statusId string) web.StatusOrderResponse
	FindAll(ctx context.Context) []web.StatusOrderResponse
}
