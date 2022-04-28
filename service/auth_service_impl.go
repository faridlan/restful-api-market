package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceImpl struct {
	UserRepository      repository.UserRepository
	BlacklistRepository repository.BlacklistRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, blacklistRepository repository.BlacklistRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return AuthServiceImpl{
		UserRepository:      userRepository,
		BlacklistRepository: blacklistRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service AuthServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		// 	ImageUrl: request.ImageUrl,
		// 	RoleId:   request.RoleId,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service AuthServiceImpl) Login(ctx context.Context, request web.LoginCreateRequest) web.Claims {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	random := helper.RandStringRunes(20)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = service.UserRepository.Login(ctx, tx, user)
	helper.PanicIfError(err)

	claim := web.Claims{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.RoleId,
		Token:    random,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(web.ExpiredTime),
		},
	}

	return helper.ToJwtResponse(claim)
}

func (service AuthServiceImpl) Profile(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service AuthServiceImpl) UpdateProfile(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Username = request.Username
	user.Email = request.Email
	user.ImageUrl = request.ImageUrl

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service AuthServiceImpl) Logout(ctx context.Context, request web.BlacklistCreateRequest) web.BlacklistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	blacklist := domain.Blacklist{
		Token: request.Token,
	}

	blacklist = service.BlacklistRepository.Create(ctx, tx, blacklist)

	return helper.ToBlacklistResponse(blacklist)
}

func (service AuthServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
