package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
)

type ShippingAddressServiceImpl struct {
	ProductRepo     repository.ProductRepository
	OrderRepo       repository.OrderRepository
	OrderDetailRepo repository.OrderDetailRepository
	CartRepo        repository.CartRepository
	DB              *sql.DB
}

func NewShippingAddressService(productRepo repository.ProductRepository, orderRepo repository.OrderRepository, orderDetailRepo repository.OrderDetailRepository, cartRepo repository.CartRepository, DB *sql.DB) ShippingAddressService {
	return ShippingAddressServiceImpl{
		ProductRepo:     productRepo,
		OrderRepo:       orderRepo,
		OrderDetailRepo: orderDetailRepo,
		CartRepo:        cartRepo,
		DB:              DB,
	}
}

func (service ShippingAddressServiceImpl) CreateOrder(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	createOrder := domain.Order{
		User: domain.User{
			Id: request.UserId,
		},
		Address: domain.Address{
			Id: request.AddressId,
		},
	}

	orders := helper.ToCreateOrders(request.Detail)

	if request.Detail[0].CartId == 0 {
		//Create Order

		createOrder = service.OrderRepo.Save(ctx, tx, createOrder)

		//Create Order Detail
		service.OrderDetailRepo.Save(ctx, tx, orders)
		//Update Total Price in Order Detail
		service.OrderDetailRepo.UpdateTotal(ctx, tx, orders)
		//Update Product Quantity
		service.OrderDetailRepo.UpdateProductQty(ctx, tx, orders)
		//Update Total in Order
		createOrder = service.OrderRepo.UpdateTotal(ctx, tx, createOrder)

		//Make Response Order
		orderDetail := service.OrderDetailRepo.FindById(ctx, tx, createOrder.Id, createOrder.User.Id)
		ordersDetail := helper.ToOrderDetailResponses(orderDetail)
		orderResponse, err := service.OrderRepo.FindById(ctx, tx, createOrder.Id, createOrder.User.Id)
		helper.PanicIfError(err)

		//Delete Cart

		return helper.ToOrderResponse(orderResponse, ordersDetail)
	} else {
		//Create Order

		createOrder = service.OrderRepo.Save(ctx, tx, createOrder)

		//Create Order Detail
		service.OrderDetailRepo.Save(ctx, tx, orders)
		//Update Total Price in Order Detail
		service.OrderDetailRepo.UpdateTotal(ctx, tx, orders)
		//Update Product Quantity
		service.OrderDetailRepo.UpdateProductQty(ctx, tx, orders)
		//Update Total in Order
		createOrder = service.OrderRepo.UpdateTotal(ctx, tx, createOrder)

		//Make Response Order
		orderDetail := service.OrderDetailRepo.FindById(ctx, tx, createOrder.Id, createOrder.User.Id)
		ordersDetail := helper.ToOrderDetailResponses(orderDetail)
		orderResponse, err := service.OrderRepo.FindById(ctx, tx, createOrder.Id, createOrder.User.Id)
		helper.PanicIfError(err)

		//Delete Cart
		carts := helper.ToDeleteOrderCarts(request.Detail)
		service.CartRepo.Delete(ctx, tx, carts)

		return helper.ToOrderResponse(orderResponse, ordersDetail)
	}

}

func (service ShippingAddressServiceImpl) FindOrderById(ctx context.Context, orderId int, userId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	orderDetail := service.OrderDetailRepo.FindById(ctx, tx, orderId, userId)
	ordersDetail := helper.ToOrderDetailResponses(orderDetail)
	order, err := service.OrderRepo.FindById(ctx, tx, orderId, userId)
	helper.PanicIfError(err)
	return helper.ToOrderResponse(order, ordersDetail)
}

func (service ShippingAddressServiceImpl) FindAllOrderByUser(ctx context.Context, userId int) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	orders, err := service.OrderRepo.FindByUserId(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToOrdersResponses(orders)
}
