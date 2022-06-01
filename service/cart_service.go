package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type ShoppingCartService interface {
	AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse
	FindCart(ctx context.Context, userId int) []web.CartResponse
	UpdateQty(ctx context.Context, request web.CartUpdateRequest) web.CartResponse
	DeleteCart(ctx context.Context, request []web.CartDeleteRequest)
	FindSomeCart(ctx context.Context, request []web.CartSelectRequest) []web.CartResponse
}
