package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type ProductDetailServiceImpl struct {
	ProductRepo repository.ProductRepository
	CartRepo    repository.CartRepository
	DB          *sql.DB
}

func NewProductDetailService(ProductRepo repository.ProductRepository, CartRepo repository.CartRepository, DB *sql.DB) ProductDetailService {
	return ProductDetailServiceImpl{
		ProductRepo: ProductRepo,
		CartRepo:    CartRepo,
		DB:          DB,
	}
}

func (service ProductDetailServiceImpl) FindProduct(ctx context.Context, productId int) web.ProductResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)
	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(product)
}

func (service ProductDetailServiceImpl) AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse {
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
	_, err = service.ProductRepo.FindById(ctx, tx, product.Product.Id)
	helper.PanicIfError(err)
	product = service.CartRepo.Save(ctx, tx, product)
	cart, err := service.CartRepo.FindById(ctx, tx, product.Product.Id)
	helper.PanicIfError(err)
	return helper.ToCartResponse(cart)
}
