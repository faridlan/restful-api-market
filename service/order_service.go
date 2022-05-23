package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ShippingAddressService interface {
	CreateOrder(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	FindOrderById(ctx context.Context, orderId int, userId int) web.OrderResponse
	FindAllOrderByUser(ctx context.Context, userId int) []web.OrderResponse
	UpdateStatus(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse
	UpdatePayment(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse
	UploadImage(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponseImg
}
