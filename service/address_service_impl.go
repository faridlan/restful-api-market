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

type AddressServiceImpl struct {
	AddressRepo repository.AddressRepository
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewAddressService(addressRepo repository.AddressRepository, DB *sql.DB, validate *validator.Validate) AddressService {
	return AddressServiceImpl{
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

	address := domain.Address{
		User: domain.User{
			Id: request.UserId,
		},
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

	address, err := service.AddressRepo.FindById(ctx, tx, request.Id, request.UserId)
	helper.PanicIfError(err)

	address.Id = request.Id
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

func (service AddressServiceImpl) Delete(ctx context.Context, addressId int, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindById(ctx, tx, addressId, userId)
	helper.PanicIfError(err)

	service.AddressRepo.Delete(ctx, tx, address)
}

func (service AddressServiceImpl) FindById(ctx context.Context, addressId int, userId int) web.AddressReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	address, err := service.AddressRepo.FindById(ctx, tx, addressId, userId)
	helper.PanicIfError(err)

	return helper.ToAddressResponse(address)
}

func (service AddressServiceImpl) FindAll(ctx context.Context, userId int) []web.AddressReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	addresses := service.AddressRepo.FindAll(ctx, tx, userId)

	return helper.ToAddressResponses(addresses)
}
