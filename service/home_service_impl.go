package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type HomeServiceImpl struct {
	ProductRepo repository.ProductRepository
	DB          *sql.DB
}

func NewHomeService(productRepo repository.ProductRepository, DB *sql.DB) HomeService {
	return HomeServiceImpl{
		ProductRepo: productRepo,
		DB:          DB,
	}
}

func (repository HomeServiceImpl) Product(ctx context.Context) []web.ProductResponse {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	productResponses := repository.ProductRepo.FindAll(ctx, tx)

	return helper.ToProductResponses(productResponses)
}
