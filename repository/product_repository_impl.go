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

	SQL := "insert into products (product_name, category_id, price, quantity, image_url) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.ProductName, product.CategoryId, product.Price, product.Quantity, product.ImageUrl)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)

	return product

}

func (repository ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	SQL := "update products set product_name = ?, category_id = ?, price = ?, quantity = ?, image_url = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ProductName, product.CategoryId, product.Price, product.Quantity, product.ImageUrl, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {

	SQL := "delete from products where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)

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
