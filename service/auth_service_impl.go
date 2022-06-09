package service

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/faridlan/restful-api-market/exception"
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
	RoleRepository      repository.RoleRepository
	Uuid                repository.UuidRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, blacklistRepository repository.BlacklistRepository, RoleRepository repository.RoleRepository, Uuid repository.UuidRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return AuthServiceImpl{
		UserRepository:      userRepository,
		BlacklistRepository: blacklistRepository,
		RoleRepository:      RoleRepository,
		Uuid:                Uuid,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service AuthServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponseLogin, web.Claims) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	stringImg := helper.NewNullString(request.ImageUrl)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	user := domain.User{
		IdUser:   uuid.Uuid,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		ImageUrl: stringImg,
		Role: domain.Role{
			Id: request.RoleId,
		},
	}

	user = service.UserRepository.Save(ctx, tx, user)
	user, err = service.UserRepository.FindById(ctx, tx, user.IdUser)
	helper.PanicIfError(err)

	random := helper.RandStringRunes(20)

	claim := web.Claims{
		Id:       user.Id,
		IdUser:   user.IdUser,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.Role.Id,
		Token:    random,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(web.ExpiredTime),
		},
	}

	claimResult := helper.ToJwtResponse(claim)
	userResult := helper.ToUserResponseLogin(user)

	return userResult, claimResult
}

func (service AuthServiceImpl) CreateUsers(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	stringImg := helper.NewNullString(request.ImageUrl)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	role, err := service.RoleRepository.FindById(ctx, tx, request.IdRole)
	helper.PanicIfError(err)

	user := domain.User{
		IdUser:   uuid.Uuid,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		ImageUrl: stringImg,
		Role: domain.Role{
			Id: role.Id,
		},
	}

	defer log.Print(user.Role.Id)

	user = service.UserRepository.SaveUsers(ctx, tx, user)
	user, err = service.UserRepository.FindById(ctx, tx, user.IdUser)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service AuthServiceImpl) Login(ctx context.Context, request web.LoginCreateRequest) (web.UserResponseLogin, web.Claims) {
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
	if err != nil {
		panic(exception.NewUnauthError(err.Error()))
	}

	claim := web.Claims{
		Id:       user.Id,
		IdUser:   user.IdUser,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.Role.Id,
		Token:    random,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(web.ExpiredTime),
		},
	}

	claimResult := helper.ToJwtResponse(claim)
	userResult := helper.ToUserResponseLogin(user)

	return userResult, claimResult
}

func (service AuthServiceImpl) Profile(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service AuthServiceImpl) UpdateProfile(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	stringImg := helper.NewNullString(request.ImageUrl)

	user, err := service.UserRepository.FindById(ctx, tx, request.IdUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.IdUser = request.IdUser
	user.Username = request.Username
	user.Email = request.Email
	user.ImageUrl = stringImg

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

func (service AuthServiceImpl) FindAll(ctx context.Context, pagination domain.Pagination) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	users := service.UserRepository.FindAll(ctx, tx, pagination)

	return helper.ToUserResponses(users)
}

func (service AuthServiceImpl) UploadImage(ctx context.Context, request web.UserCreateRequest) web.UserResponseImg {
	random := helper.RandStringRunes(10)
	s3Client, endpoint := helper.S3Config()

	object := s3.PutObjectInput{
		Bucket: aws.String("olshop"),
		Key:    aws.String("/profiles/" + random + ".png"),
		Body:   strings.NewReader(string(request.ImageUrl)),
		ACL:    aws.String("public-read"),
	}

	_, err := s3Client.PutObject(&object)
	helper.PanicIfError(err)

	image := web.UserResponseImg{
		ImageUrl: "https://" + *object.Bucket + "." + endpoint + *object.Key,
	}

	return image
}

func (service AuthServiceImpl) FindSeeder(ctx context.Context, pagination domain.Pagination) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	users, err := service.UserRepository.FindSeeder(ctx, tx, pagination)
	helper.PanicIfError(err)

	return helper.ToUserResponse(users)
}
