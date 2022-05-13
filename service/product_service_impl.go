package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
}

func NewProductServie(productRepository repository.ProductRepository, DB *sql.DB) ProductService {
	return ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product := domain.Product{
		ProductName: request.ProductName,
		CategoryId:  request.CategoryId,
		Price:       request.Price,
		Quantity:    request.Quantity,
		ImageUrl:    request.ImageUrl,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Update(ctx context.Context, request web.ProductUpateRequest) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	product.ProductName = request.ProductName
	product.CategoryId = request.CategoryId
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.ImageUrl = request.ImageUrl

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	service.ProductRepository.Delete(ctx, tx, product)
}

func (repository ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	productResponses := repository.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(productResponses)
}
