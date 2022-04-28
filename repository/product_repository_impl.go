package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProdcutRepository() ProductRepository {
	return ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	panic("not implemented") // TODO: Implement
}

func (repository ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	panic("not implemented") // TODO: Implement
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	panic("not implemented") // TODO: Implement
}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "select id, product_name, category_id, price, quantity, image_url from products where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()
	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.ProductName, &product.CategoryId, &product.Price, &product.Quantity, &product.ImageUrl)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {

	SQL := "select id, product_name, category_id, price, image_url from products limit 10"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.ProductName, &product.CategoryId, &product.Price, &product.ImageUrl)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
