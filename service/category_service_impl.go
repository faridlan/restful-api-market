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

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepository
	Uuid         repository.UuidRepository
	Db           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, Uuid repository.UuidRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return CategoryServiceImpl{
		CategoryRepo: categoryRepo,
		Db:           db,
		Uuid:         Uuid,
		Validate:     validate,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	uuid, err := service.Uuid.CreteUui(ctx, tx)
	helper.PanicIfError(err)

	category := domain.Category{
		CategoryName: request.CategoryName,
		IdCategory:   uuid.Uuid,
	}

	category = service.CategoryRepo.Save(ctx, tx, category)
	category, err = service.CategoryRepo.FindById(ctx, tx, category.IdCategory)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, request.IdCategory)
	helper.PanicIfError(err)

	category.IdCategory = request.IdCategory
	category.CategoryName = request.CategoryName

	category = service.CategoryRepo.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepo.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	categories := service.CategoryRepo.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
