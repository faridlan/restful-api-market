package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ShippingAddressService interface {
	Order(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	FindOrderById(ctx context.Context, orderId int, userId int) web.OrderResponse
	FindAllOrderByUser(ctx context.Context, userId int) []web.OrderResponse
}
