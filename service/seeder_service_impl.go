package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/repository"
)

type SeederServiceImpl struct {
	AddressRepo     repository.AddressRepository
	OrderDetailRepo repository.OrderDetailRepository
	OrderRepo       repository.OrderRepository
	ProductRepo     repository.ProductRepository
	UserRepo        repository.UserRepository
	DB              *sql.DB
}

func NewSeedService(AddressRepo repository.AddressRepository, OrderDetailRepo repository.OrderDetailRepository, OrderRepo repository.OrderRepository, ProductRepo repository.ProductRepository, UserRepo repository.UserRepository, DB *sql.DB) SeederServiceImpl {
	return SeederServiceImpl{
		AddressRepo:     AddressRepo,
		OrderDetailRepo: OrderDetailRepo,
		OrderRepo:       OrderRepo,
		ProductRepo:     ProductRepo,
		UserRepo:        UserRepo,
		DB:              DB,
	}
}

func (service SeederServiceImpl) Delete(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	service.OrderDetailRepo.DeleteTable(ctx, tx)
	service.OrderRepo.DeleteTable(ctx, tx)
	service.ProductRepo.DeleteTable(ctx, tx)
	service.AddressRepo.DeleteTable(ctx, tx)
	service.UserRepo.DeleteTable(ctx, tx)
}
