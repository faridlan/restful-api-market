package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
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

func (service ShoppingCartServiceImpl) AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)
	product := domain.Cart{
		User: domain.User{
			Id: request.UserId,
		},
		Product: domain.Product{
			Id: request.ProductId,
		},
		Quantity: request.Quantity,
	}

	product = service.CartRepository.Save(ctx, tx, product)
	cart, err := service.CartRepository.FindById(ctx, tx, product.Product.Id)
	helper.PanicIfError(err)
	return helper.ToCartResponse(cart)
}

func (service ShoppingCartServiceImpl) FindCart(ctx context.Context, userId int) []web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)

	cartResponses, err := service.CartRepository.FindAll(ctx, tx, userId)
	helper.PanicIfError(err)

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

	carts := helper.ToCartsDelete(request)
	service.CartRepository.Delete(ctx, tx, carts)
}
