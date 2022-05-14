package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, request web.PaymentCreateRequest) web.PaymentResponse
	UpdateOrder(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse
}
