package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponseLogin, web.Claims)
	CreateUsers(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.LoginCreateRequest) (web.UserResponseLogin, web.Claims)
	Profile(ctx context.Context, userId string) web.UserResponse
	UpdateProfile(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	FindAll(ctx context.Context, pagination domain.Pagination) []web.UserResponse
	Logout(ctx context.Context, request web.BlacklistCreateRequest) web.BlacklistResponse
	UploadImage(ctx context.Context, storage domain.Storage) web.UserResponseImg
	FindSeeder(ctx context.Context, pagination domain.Pagination) web.UserResponse
}
