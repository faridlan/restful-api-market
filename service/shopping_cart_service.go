package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ShoppingCartService interface {
	FindCart(ctx context.Context, userId int) []web.CartResponse
	UpdateQty(ctx context.Context, request web.CartUpdateRequest) web.CartResponse
	DeleteCart(ctx context.Context, request []web.CartDeleteRequest)
}
