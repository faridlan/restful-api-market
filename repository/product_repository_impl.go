package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProdcutRepository() ProductRepository {
	return ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into products (id_product, product_name, category_id, price, quantity, image_url) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.IdProduct, product.ProductName, product.Category.Id, product.Price, product.Quantity, product.ImageUrl)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)

	return product
}

func (repository ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	SQL := "update products set product_name = ?, category_id = ?, price = ?, quantity = ?, image_url = ? where id_product = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ProductName, product.Category.Id, product.Price, product.Quantity, product.ImageUrl, product.IdProduct)
	helper.PanicIfError(err)

	return product
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {

	SQL := "delete from products where id_product = ?"
	_, err := tx.ExecContext(ctx, SQL, product.IdProduct)
	helper.PanicIfError(err)

}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	SQL := `select p.id, p.id_product, p.product_name, c.id_category, c.category_name, p.price, p.quantity, p.image_url 
	from products as p
	inner join categories as c on c.id = p.category_id
	where p.id_product = ?`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()
	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.IdProduct, &product.ProductName, &product.Category.IdCategory, &product.Category.CategoryName, &product.Price, &product.Quantity, &product.ImageUrl)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) []domain.Product {

	SQL := fmt.Sprintf(`select p.id_product, p.product_name, c.id_category, c.category_name, p.price, p.quantity, p.image_url 
	from products as p
	inner join categories as c on c.id = p.category_id
	where p.quantity > 0
	order by p.id desc limit %d,%d`, pagination.Page, pagination.Limit)
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.IdProduct, &product.ProductName, &product.Category.IdCategory, &product.Category.CategoryName, &product.Price, &product.Quantity, &product.ImageUrl)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}

func (repository ProductRepositoryImpl) FindSeeder(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) []domain.Product {
	SQL := fmt.Sprintf("select id, id_product from products order by id limit %d,%d", pagination.Page, pagination.Limit)
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.IdProduct)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (repository ProductRepositoryImpl) DeleteTable(ctx context.Context, tx *sql.Tx) {
	SQL := "delete from products"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
