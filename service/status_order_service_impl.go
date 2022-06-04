package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/exception"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/go-playground/validator/v10"
)

type StatusOrderServiceImpl struct {
	Repository repository.StatusOrderRepository
	Uuid       repository.UuidRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewStatusOrderService(repository repository.StatusOrderRepository, Uuid repository.UuidRepository, db *sql.DB, validate *validator.Validate) StatusOrderService {
	return StatusOrderServiceImpl{
		Repository: repository,
		Uuid:       Uuid,
		DB:         db,
		Validate:   validate,
	}
}

func (service StatusOrderServiceImpl) Create(ctx context.Context, request web.StatusOrderCreate) web.StatusOrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	statusOrder := domain.StatusOrder{
		IdStatusOrder: uuid.Uuid,
		StatusName:    request.StatusName,
	}
	statusOrder = service.Repository.Save(ctx, tx, statusOrder)

	return helper.ToStatusOrderResponse(statusOrder)

}

func (service StatusOrderServiceImpl) Update(ctx context.Context, request web.StatusOrderUpdate) web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, request.IdStatusOrder)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	statusOrder.IdStatusOrder = request.IdStatusOrder
	statusOrder.StatusName = request.StatusName
	statusOrder = service.Repository.Update(ctx, tx, statusOrder)

	return helper.ToStatusOrderResponse(statusOrder)
}

func (service StatusOrderServiceImpl) Delete(ctx context.Context, statusId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, statusId)
	helper.PanicIfError(err)

	service.Repository.Delete(ctx, tx, statusOrder)
}

func (service StatusOrderServiceImpl) FindById(ctx context.Context, statusId string) web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrder, err := service.Repository.FindById(ctx, tx, statusId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStatusOrderResponse(statusOrder)
}

func (service StatusOrderServiceImpl) FindAll(ctx context.Context) []web.StatusOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	statusOrders := service.Repository.FindAll(ctx, tx)

	return helper.ToStatusOrderResponses(statusOrders)
}
