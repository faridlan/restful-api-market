package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.LoginCreateRequest) web.Claims
	Profile(ctx context.Context, userId int) web.UserResponse
	UpdateProfile(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	Logout(ctx context.Context, request web.BlacklistCreateRequest) web.BlacklistResponse
	// ChangePassword(ctx context.Context, request)
}
