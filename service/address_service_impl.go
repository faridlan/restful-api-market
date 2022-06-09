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

type AddressServiceImpl struct {
	AddressRepo repository.AddressRepository
	Uuid        repository.UuidRepository
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewAddressService(addressRepo repository.AddressRepository, Uuid repository.UuidRepository, DB *sql.DB, validate *validator.Validate) AddressService {
	return AddressServiceImpl{
		Uuid:        Uuid,
		AddressRepo: addressRepo,
		DB:          DB,
		Validate:    validate,
	}
}

func (service AddressServiceImpl) Create(ctx context.Context, request web.AddressCreateRequest) web.AddressReponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)
	address := domain.Address{
		IdAddress:       uuid.Uuid,
		User:            domain.User{Id: request.UserId},
		Name:            request.Name,
		HandphoneNumber: request.HandphoneNumber,
		Street:          request.Street,
		Districk:        request.Districk,
		PostCode:        request.PostCode,
		Comment:         request.Comment,
	}

	address = service.AddressRepo.Save(ctx, tx, address)

	return helper.ToAddressResponse(address)
}

func (service AddressServiceImpl) Update(ctx context.Context, request web.AddressUpdateRequest) web.AddressReponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindById(ctx, tx, request.IdAddress, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	address.Id = request.Id
	address.IdAddress = request.IdAddress
	address.User.Id = request.UserId
	address.Name = request.Name
	address.HandphoneNumber = request.HandphoneNumber
	address.Street = request.Street
	address.Districk = request.Districk
	address.PostCode = request.PostCode
	address.Comment = request.Comment

	address = service.AddressRepo.Update(ctx, tx, address)

	return helper.ToAddressResponse(address)
}

func (service AddressServiceImpl) Delete(ctx context.Context, addressId string, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindById(ctx, tx, addressId, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AddressRepo.Delete(ctx, tx, address)
}

func (service AddressServiceImpl) FindById(ctx context.Context, addressId string, userId int) web.AddressReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindById(ctx, tx, addressId, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAddressResponse(address)
}

func (service AddressServiceImpl) FindAll(ctx context.Context, userId int, pagination domain.Pagination) []web.AddressReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	addresses, err := service.AddressRepo.FindAll(ctx, tx, userId, pagination)
	helper.PanicIfError(err)

	return helper.ToAddressResponses(addresses)
}

func (service AddressServiceImpl) FindSeeder(ctx context.Context, pagination domain.Pagination) web.AddressReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindSeeder(ctx, tx, pagination)
	helper.PanicIfError(err)

	return helper.ToAddressResponse(address)
}
