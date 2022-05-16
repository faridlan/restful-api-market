package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductServie(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product := domain.Product{
		ProductName: request.ProductName,
		Category: domain.Category{
			Id: request.CategoryId,
		},
		Price:    request.Price,
		Quantity: request.Quantity,
		ImageUrl: request.ImageUrl,
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
	product.Category.Id = request.CategoryId
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

func (service ProductServiceImpl) FindyId(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	productResponses := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(productResponses)
}
