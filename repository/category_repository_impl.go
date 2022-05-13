package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "insert into categories(category_name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.CategoryName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "update categories set category_name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryName, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {

	SQL := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, category_name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.CategoryName)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, category_name from categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.CategoryName)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
