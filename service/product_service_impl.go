package service

import (
	"context"
	"database/sql"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Uuid              repository.UuidRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductServie(productRepository repository.ProductRepository, Uuid repository.UuidRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return ProductServiceImpl{
		ProductRepository: productRepository,
		Uuid:              Uuid,
		DB:                DB,
		Validate:          validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	product := domain.Product{
		IdProduct:   uuid.Uuid,
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

	product, err := service.ProductRepository.FindById(ctx, tx, request.IdProduct)
	helper.PanicIfError(err)

	product.IdProduct = request.IdProduct
	product.ProductName = request.ProductName
	product.Category.Id = request.CategoryId
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.ImageUrl = request.ImageUrl

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service ProductServiceImpl) FindyId(ctx context.Context, productId string) web.ProductResponse {
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

func (servicer ProductServiceImpl) CreateImg(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {

	random := helper.RandStringRunes(10)
	s3Client, endpoint := helper.S3Config()

	object := s3.PutObjectInput{
		Bucket: aws.String("olshop"),
		Key:    aws.String("/products/" + random + ".png"),
		Body:   strings.NewReader(string(request.ImageUrl)),
		ACL:    aws.String("public-read"),
	}

	_, err := s3Client.PutObject(&object)
	helper.PanicIfError(err)

	image := web.ProductResponse{
		ImageUrl: "https://" + *object.Bucket + "." + endpoint + *object.Key,
	}

	return image

}
