package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type PaymentServiceImpl struct {
	OrderRepo   repository.OrderRepository
	PaymentRepo repository.PaymentRepository
	DB          *sql.DB
}

func (service PaymentServiceImpl) CreatePayment(ctx context.Context, request web.PaymentCreateRequest) web.PaymentResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	payment := domain.Payment{
		ImageUrl: request.ImageUrl,
	}

	payment = service.PaymentRepo.Save(ctx, tx, payment)

	order := domain.Order{
		Id: request.OrdeId,
		// PaymentId: payment.Id,
	}

	order = service.OrderRepo.UpdatePayment(ctx, tx, order)

	return web.PaymentResponse{}
}

func (service PaymentServiceImpl) UpdateOrder(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	panic("not implemented") // TODO: Implement
}
