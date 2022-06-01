package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return OrderRepositoryImpl{}
}

func (repository OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "insert into orders (user_id, address_id) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, order.User.Id, order.Address.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	order.Id = int(id)

	return order
}

func (repository OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int, userId int) (domain.Order, error) {
	SQL := `select o.id, u.username,
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date, s.status_name, o.payment from orders as o
	inner join users as u on u.id = o.user_id
	inner join addresses as a on a.id = o.address_id
	inner join status_order as s on s.id = o.status_id
	where o.id = ? and o.user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, orderId, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.User.Username, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order not found")
	}
}

func (repository OrderRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Order, error) {
	SQL := `select o.id, u.username, 
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date from orders as o
	inner join users as u on u.id = o.user_id
	inner join addresses as a on a.id = o.address_id
	where o.user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	orders := []domain.Order{}

	if rows.Next() {
		for rows.Next() {
			order := domain.Order{}
			err := rows.Scan(&order.Id, &order.User.Username, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate)
			helper.PanicIfError(err)

			orders = append(orders, order)
		}
		return orders, nil
	} else {
		return orders, errors.New("order not found")
	}
}

func (repository OrderRepositoryImpl) UpdateTotal(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	// SQL := `UPDATE orders
	// INNER JOIN orders_detail ON (orders.id = orders_detail.order_id)
	// SET orders.total = (select sum(total_price) from orders_detail where order_id = (select max(orders.id) where orders.user_id = ?))
	// WHERE orders_detail.order_id = (select max(orders.id) where orders.user_id = ?)`

	SQL := `update orders set total = (select sum(total_price) from orders_detail where order_id = ?)
	where id = ? and user_id = ?`

	_, err := tx.ExecContext(ctx, SQL, order.Id, order.Id, order.User.Id)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set status_id = ? where id = ? and user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Status.Id, order.Id, order.User.Id)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) UpdatePayment(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set payment = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Payment, order.Id)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	SQL := `select o.id, u.username, 
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date from orders as o
	inner join users as u on u.id = o.user_id
	inner join addresses as a on a.id = o.address_id`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	orders := []domain.Order{}

	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.Id, &order.User.Username, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate)
		helper.PanicIfError(err)

		orders = append(orders, order)
	}
	return orders
}

func (repository OrderRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, orderId int) (domain.Order, error) {
	SQL := `select o.id, u.username,
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date, s.status_name, o.payment from orders as o
	inner join users as u on u.id = o.user_id
	inner join addresses as a on a.id = o.address_id
	inner join status_order as s on s.id = o.status_id
	where o.id = ?`

	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError(err)

	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.User.Username, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order not found")
	}
}
