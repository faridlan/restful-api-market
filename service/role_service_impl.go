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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Uuid           repository.UuidRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, Uuid repository.UuidRepository, DB *sql.DB, validate *validator.Validate) RoleService {
	return RoleServiceImpl{
		RoleRepository: roleRepository,
		Uuid:           Uuid,
		DB:             DB,
		Validate:       validate,
	}
}

func (service RoleServiceImpl) Create(ctx context.Context, request web.RoleCreateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	role := domain.Role{
		IdRole: uuid.Uuid,
		Name:   request.Name,
	}

	role = service.RoleRepository.Save(ctx, tx, role)
	role, err = service.RoleRepository.FindById(ctx, tx, role.IdRole)
	helper.PanicIfError(err)

	return helper.ToRoleResponse(role)
}

func (service RoleServiceImpl) Update(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, request.IdRole)
	helper.PanicIfError(err)

	role.IdRole = request.IdRole
	role.Name = request.Name

	role = service.RoleRepository.Update(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service RoleServiceImpl) FindById(ctx context.Context, roleId string) web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	helper.PanicIfError(err)

	return helper.ToRoleResponse(role)
}

func (service RoleServiceImpl) FindAll(ctx context.Context) []web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	roles := service.RoleRepository.FindAll(ctx, tx)

	return helper.ToRoleResponses(roles)
}
