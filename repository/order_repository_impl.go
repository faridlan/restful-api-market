package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return OrderRepositoryImpl{}
}

func (repository OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "insert into orders (id_order, user_id, address_id) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, order.IdOrder, order.User.Id, order.Address.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	order.Id = int(id)

	return order
}

func (repository OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId string, userId int) (domain.Order, error) {
	SQL := `select o.id, o.id_order, a.id_address,
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date, s.id_status_order, s.status_name, o.payment from orders as o
	inner join addresses as a on a.id = o.address_id
	inner join status_order as s on s.id = o.status_id
	where o.id_order = ? and o.user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, orderId, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.IdOrder, &order.Address.IdAddress, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate, &order.Status.IdStatusOrder, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order not found")
	}
}

func (repository OrderRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int, pagination domain.Pagination) ([]domain.Order, error) {
	SQL := fmt.Sprintf(`select o.id_order,
	o.total, o.order_date, s.id_status_order, s.status_name, o.payment from orders as o
	inner join status_order as s on s.id = o.status_id
	where o.user_id = ?
	order by o.id desc limit %d,%d`, pagination.Page, pagination.Limit)

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	orders := []domain.Order{}
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.IdOrder, &order.Total, &order.OrderDate, &order.Status.IdStatusOrder, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		orders = append(orders, order)
	}
	return orders, nil
}

func (repository OrderRepositoryImpl) UpdateTotal(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := `update orders set total = (select sum(total_price) from orders_detail where order_id = ?)
	where id = ? and user_id = ?`

	_, err := tx.ExecContext(ctx, SQL, order.Id, order.Id, order.User.Id)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set status_id = ? where id_order = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Status.Id, order.IdOrder)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) UpdatePayment(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set payment = ? where id_order = ? and user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Payment, order.IdOrder, order.User.Id)
	helper.PanicIfError(err)

	return order
}

func (repository OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) []domain.Order {
	SQL := fmt.Sprintf(`select o.id_order,
	o.total, o.order_date, s.id_status_order, s.status_name, o.payment from orders as o
	inner join status_order as s on s.id = o.status_id
	order by o.payment desc limit %d,%d`, pagination.Page, pagination.Limit)

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	orders := []domain.Order{}

	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.IdOrder, &order.Total, &order.OrderDate, &order.Status.IdStatusOrder, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		orders = append(orders, order)
	}
	return orders
}

func (repository OrderRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error) {
	SQL := `select o.id, o.id_order, a.id_address,
	a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment,
	o.total, o.order_date, s.id_status_order, s.status_name, o.payment from orders as o
	inner join addresses as a on a.id = o.address_id
	inner join status_order as s on s.id = o.status_id
	where o.id_order = ?`

	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError(err)

	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.IdOrder, &order.Address.IdAddress, &order.Address.Name, &order.Address.HandphoneNumber, &order.Address.Street, &order.Address.Districk, &order.Address.PostCode, &order.Address.Comment, &order.Total, &order.OrderDate, &order.Status.IdStatusOrder, &order.Status.StatusName, &order.Payment)
		helper.PanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order not found")
	}
}

func (repository OrderRepositoryImpl) DeleteTable(ctx context.Context, tx *sql.Tx) {
	SQL := "delete from orders"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
