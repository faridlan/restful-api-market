package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type ShoppingCartServiceImpl struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
}

func NewShoppingCartService(CartRepository repository.CartRepository, DB *sql.DB) ShoppingCartService {
	return ShoppingCartServiceImpl{
		CartRepository: CartRepository,
		DB:             DB,
	}
}

func (service ShoppingCartServiceImpl) FindCart(ctx context.Context, userId int) []web.CartResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)

	cartResponses := service.CartRepository.FindAll(ctx, tx, userId)

	return helper.ToCartResponses(cartResponses)
}

func (service ShoppingCartServiceImpl) UpdateQty(ctx context.Context, request web.CartUpdateRequest) web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)

	cart, err := service.CartRepository.FindById(ctx, tx, request.ProductId)
	helper.PanicIfError(err)

	cart.User.Id = request.UserId
	cart.Product.Id = request.ProductId
	cart.Quantity = request.Quantity

	cart = service.CartRepository.Update(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service ShoppingCartServiceImpl) DeleteCart(ctx context.Context, request []web.CartDeleteRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)

	// cart, err := service.CartRepository.FindById(ctx, tx, productId)
	// helper.PanicIfError(err)

	carts := helper.ToCartsDelete(request)
	service.CartRepository.Delete(ctx, tx, carts)
}
