package service

import (
	"context"
	"database/sql"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/faridlan/restful-api-market/exception"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/go-playground/validator/v10"
)

type ShippingAddressServiceImpl struct {
	ProductRepo     repository.ProductRepository
	OrderRepo       repository.OrderRepository
	OrderDetailRepo repository.OrderDetailRepository
	CartRepo        repository.CartRepository
	AddressRepo     repository.AddressRepository
	StatusRepo      repository.StatusOrderRepository
	Uuid            repository.UuidRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewShippingAddressService(productRepo repository.ProductRepository, orderRepo repository.OrderRepository, orderDetailRepo repository.OrderDetailRepository, cartRepo repository.CartRepository, AddressRepo repository.AddressRepository, StatusRepo repository.StatusOrderRepository, Uuid repository.UuidRepository, DB *sql.DB, validate *validator.Validate) ShippingAddressService {
	return ShippingAddressServiceImpl{
		ProductRepo:     productRepo,
		OrderRepo:       orderRepo,
		OrderDetailRepo: orderDetailRepo,
		CartRepo:        cartRepo,
		AddressRepo:     AddressRepo,
		StatusRepo:      StatusRepo,
		Uuid:            Uuid,
		DB:              DB,
		Validate:        validate,
	}
}

func (service ShippingAddressServiceImpl) CreateOrder(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	address, err := service.AddressRepo.FindById(ctx, tx, request.AddressId, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	createOrder := domain.Order{
		IdOrder: uuid.Uuid,
		User: domain.User{
			Id: request.UserId,
		},
		Address: domain.Address{
			Id: address.Id,
		},
	}

	x := []domain.Product{}
	for _, v := range request.Products {
		product, err := service.ProductRepo.FindById(ctx, tx, v.ProductId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		x = append(x, product)
	}

	orders := helper.ToCreateOrders(request.Products)
	ordersCreate := []domain.OrderDetail{}
	for i, v := range x {
		z := orders[i]
		z.Product.Id = v.Id
		ordersCreate = append(ordersCreate, z)
	}

	//Create Order

	createOrder = service.OrderRepo.Save(ctx, tx, createOrder)

	//Create Order Detail
	service.OrderDetailRepo.Save(ctx, tx, ordersCreate)
	//Update Total Price in Order Detail
	service.OrderDetailRepo.UpdateTotal(ctx, tx, ordersCreate)
	//Update Product Quantity
	service.OrderDetailRepo.UpdateProductQty(ctx, tx, ordersCreate)
	//Update Total in Order
	createOrder = service.OrderRepo.UpdateTotal(ctx, tx, createOrder)

	//Make Response Order
	orderDetail := service.OrderDetailRepo.FindById(ctx, tx, createOrder.Id, createOrder.User.Id)
	ordersDetail := helper.ToOrderDetailResponses(orderDetail)
	orderResponse, err := service.OrderRepo.FindById(ctx, tx, createOrder.IdOrder, createOrder.User.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//Delete Cart

	return helper.ToOrderResponse(orderResponse, ordersDetail)

}

func (service ShippingAddressServiceImpl) FindOrderById(ctx context.Context, orderId string, userId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	order, err := service.OrderRepo.FindById(ctx, tx, orderId, userId)
	helper.PanicIfError(err)
	orderDetail := service.OrderDetailRepo.FindById(ctx, tx, order.Id, userId)
	ordersDetail := helper.ToOrderDetailResponses(orderDetail)
	return helper.ToOrderResponse(order, ordersDetail)
}

func (service ShippingAddressServiceImpl) FindAllOrderByUser(ctx context.Context, userId int, pagination domain.Pagination) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	orders, err := service.OrderRepo.FindByUserId(ctx, tx, userId, pagination)
	helper.PanicIfError(err)

	return helper.ToOrdersResponses(orders)
}

func (service ShippingAddressServiceImpl) UpdateStatus(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	order, err := service.OrderRepo.FindId(ctx, tx, request.IdOrder)
	helper.PanicIfError(err)
	status, err := service.StatusRepo.FindById(ctx, tx, request.IdStatus)
	helper.PanicIfError(err)

	order.Status.Id = status.Id
	order.IdOrder = request.IdOrder
	order.User.Id = request.UserId

	order = service.OrderRepo.UpdateStatus(ctx, tx, order)
	order, err = service.OrderRepo.FindId(ctx, tx, request.IdOrder)
	helper.PanicIfError(err)
	// orderDetail := service.OrderDetailRepo.FindById(ctx, tx, request.OrderId, request.UserId)
	// ordersDetail := helper.ToOrderDetailResponses(orderDetail)

	return helper.ToOrdersResponse(order)
}

func (service ShippingAddressServiceImpl) UpdatePayment(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	stringX := helper.NewNullString(request.Payment)

	orderDetail := service.OrderDetailRepo.FindById(ctx, tx, request.OrderId, request.UserId)
	ordersDetail := helper.ToOrderDetailResponses(orderDetail)
	order := domain.Order{}

	order.Status.Id = request.StatusId
	order.Payment = stringX
	// order.Payment = request.Payment
	order.IdOrder = request.IdOrder
	order.User.Id = request.UserId

	order = service.OrderRepo.UpdatePayment(ctx, tx, order)
	order = service.OrderRepo.UpdateStatus(ctx, tx, order)

	order, err = service.OrderRepo.FindById(ctx, tx, request.IdOrder, request.UserId)
	helper.PanicIfError(err)
	return helper.ToOrderResponse(order, ordersDetail)

}

func (service ShippingAddressServiceImpl) UploadImage(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponseImg {
	random := helper.RandStringRunes(10)
	s3Client, endpoint := helper.S3Config()

	object := s3.PutObjectInput{
		Bucket: aws.String("olshop"),
		Key:    aws.String("/payments/" + random + ".png"),
		Body:   strings.NewReader(string(request.Payment)),
		ACL:    aws.String("public-read"),
	}

	_, err := s3Client.PutObject(&object)
	helper.PanicIfError(err)

	image := web.OrderResponseImg{
		Image: "https://" + *object.Bucket + "." + endpoint + *object.Key,
	}

	return image
}

func (service ShippingAddressServiceImpl) FindAll(ctx context.Context, pagination domain.Pagination) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	orders := service.OrderRepo.FindAll(ctx, tx, pagination)

	return helper.ToOrdersResponses(orders)
}

func (service ShippingAddressServiceImpl) FindById(ctx context.Context, orderId string) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	order, err := service.OrderRepo.FindId(ctx, tx, orderId)
	helper.PanicIfError(err)
	orders := service.OrderDetailRepo.AdminFindById(ctx, tx, order.Id)
	ordersDetail := helper.ToOrderDetailResponses(orders)

	return helper.ToOrderResponse(order, ordersDetail)
}
