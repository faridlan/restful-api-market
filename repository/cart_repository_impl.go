package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepository() CartRepository {
	return CartRepositoryImpl{}
}

func (repository CartRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into carts (user_id, product_id, quantity) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, cart.User.Id, cart.Product.Id, cart.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cart.Id = int(id)

	return cart
}

func (repository CartRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "update carts set quantity = ? where product_id = ? and user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Quantity, cart.Product.Id, cart.User.Id)
	helper.PanicIfError(err)

	return cart
}

func (repository CartRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, carts []domain.Cart) {
	SQL := "delete from carts where user_id = ? and product_id = ?"
	for _, cart := range carts {
		_, err := tx.ExecContext(ctx, SQL, cart.User.Id, cart.Product.Id)
		helper.PanicIfError(err)
	}
}

func (repository CartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Cart, error) {
	SQL := `select c.id, u.id, u.username, c.quantity, p.id, p.product_name, ct.category_name, p.price, p.quantity, p.image_url from carts as c
	left join products as p on p.id = c.product_id
	left join categories as ct on ct.id = p.category_id
	left join users as u on u.id = c.user_id
	where c.product_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	cart := domain.Cart{}

	if rows.Next() {
		err := rows.Scan(&cart.Id, &cart.User.Id, &cart.User.Username, &cart.Quantity, &cart.Product.Id, &cart.Product.ProductName, &cart.Product.Category.CategoryName, &cart.Product.Price, &cart.Product.Quantity, &cart.Product.ImageUrl)
		helper.PanicIfError(err)

		return cart, nil
	} else {
		return cart, errors.New("cart not found")
	}
}

func (repository CartRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Cart, error) {
	SQL := `select c.id, u.id, u.username, c.quantity, p.id, p.product_name, ct.category_name, p.price, p.quantity, p.image_url from carts as c
	left join products as p on p.id = c.product_id
	left join categories as ct on ct.id = p.category_id
	left join users as u on u.id = c.user_id
	where c.user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	carts := []domain.Cart{}

	for rows.Next() {
		cart := domain.Cart{}
		if err = rows.Scan(&cart.Id, &cart.User.Id, &cart.User.Username, &cart.Quantity, &cart.Product.Id, &cart.Product.ProductName, &cart.Product.Category.CategoryName, &cart.Product.Price, &cart.Product.Quantity, &cart.Product.ImageUrl); err != nil {
			return carts, err
		}

		carts = append(carts, cart)
	}
	if err = rows.Err(); err != nil {
		return carts, err
	}

	return carts, nil
}
