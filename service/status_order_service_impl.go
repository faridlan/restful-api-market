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

type StatusOrderServiceImpl struct {
	Repository repository.StatusOrderRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewStatusOrderService(repository repository.StatusOrderRepository, db *sql.DB, validate *validator.Validate) StatusOrderService {
	return StatusOrderServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service StatusOrderServiceImpl) Create(ctx context.Context, request web.StatusOrderCreate) web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder := domain.StatusOrder{
		StatusName: request.StatusName,
	}
	statusOrder = service.Repository.Save(ctx, tx, statusOrder)

	return helper.ToStatusOrderResponse(statusOrder)

}

func (service StatusOrderServiceImpl) Update(ctx context.Context, request web.StatusOrderUpdate) web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	statusOrder.Id = request.Id
	statusOrder.StatusName = request.StatusName
	statusOrder = service.Repository.Update(ctx, tx, statusOrder)

	return helper.ToStatusOrderResponse(statusOrder)
}

func (service StatusOrderServiceImpl) Delete(ctx context.Context, statusId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, statusId)
	helper.PanicIfError(err)

	service.Repository.Delete(ctx, tx, statusOrder)
}

func (service StatusOrderServiceImpl) FindById(ctx context.Context, statusId int) web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, statusId)
	helper.PanicIfError(err)

	return helper.ToStatusOrderResponse(statusOrder)
}

func (service StatusOrderServiceImpl) FindAll(ctx context.Context) []web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrders := service.Repository.FindAll(ctx, tx)

	return helper.ToStatusOrderResponses(statusOrders)
}
