package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type OrderDetailRepositoryImpl struct {
}

func NewOrderDetailRepository() OrderDetailRepository {
	return OrderDetailRepositoryImpl{}
}

func (repository OrderDetailRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orders []domain.OrderDetail) []domain.OrderDetail {
	SqlScript := "insert into orders_detail (order_id, product_id, quantity, unit_price) values"
	var vals []interface{}

	for _, order := range orders {
		SqlScript += "((select max(id) from orders where user_id = ?), ?, ?, (select price from products where id = ?)),"
		vals = append(vals, order.Order.User.Id, order.Product.Id, order.Quantity, order.Product.Id)
	}

	SqlScript = SqlScript[0 : len(SqlScript)-1]

	Statement, err := tx.PrepareContext(ctx, SqlScript)
	if err != nil {
		panic(err)
	}
	defer Statement.Close()

	_, err = Statement.Exec(vals...)
	if err != nil {
		panic(err)
	}

	return orders

}
func (repository OrderDetailRepositoryImpl) UpdateTotal(ctx context.Context, tx *sql.Tx, orders []domain.OrderDetail) []domain.OrderDetail {
	SQL := `update orders_detail set total_price = (select unit_price * quantity where product_id = ? and order_id = (select max(id) from orders where user_id = ?)) 
	where product_id = ? and order_id = (select max(id) from orders where user_id = ?)`

	for _, order := range orders {
		_, err := tx.ExecContext(ctx, SQL, order.Product.Id, order.Order.User.Id, order.Product.Id, order.Order.User.Id)
		if err != nil {
			panic(err)
		}
	}

	return orders
}

func (repository OrderDetailRepositoryImpl) UpdateProductQty(ctx context.Context, tx *sql.Tx, products []domain.OrderDetail) []domain.OrderDetail {
	SQL := `update products
	set quantity = quantity - (select quantity from orders_detail where product_id = ?
	and order_id = (select max(id) from orders where user_id = ?)) where id = ?;`

	for _, product := range products {
		_, err := tx.ExecContext(ctx, SQL, product.Product.Id, product.Order.User.Id, product.Product.Id)
		helper.PanicIfError(err)
	}

	return products
}

func (repository OrderDetailRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int, userId int) []domain.OrderDetail {
	SQL := `select p.product_name, d.quantity, d.unit_price, d.total_price from orders as o
	left join orders_detail as d on d.order_id = o.id
	left join products as p on p.id = d.product_id 
	where d.order_id = ? and user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, orderId, userId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	orders := []domain.OrderDetail{}

	for rows.Next() {
		order := domain.OrderDetail{}
		err := rows.Scan(&order.Product.ProductName, &order.Quantity, &order.UnitPrice, &order.TotalPrice)
		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	return orders
}
