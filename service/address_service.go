package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
)

type AddressService interface {
	Create(ctx context.Context, request web.AddressCreateRequest) web.AddressReponse
	Update(ctx context.Context, request web.AddressUpdateRequest) web.AddressReponse
	Delete(ctx context.Context, addressId string, userId int)
	FindById(ctx context.Context, addressId string, userId int) web.AddressReponse
	FindAll(ctx context.Context, userId int, pagination domain.Pagination) []web.AddressReponse
	FindSeeder(ctx context.Context, pagination domain.Pagination) web.AddressReponse
}
