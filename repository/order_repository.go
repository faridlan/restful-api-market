package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	FindById(ctx context.Context, tx *sql.Tx, orderId string, userId int) (domain.Order, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int, pagination domain.Pagination) ([]domain.Order, error)
	UpdateTotal(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	UpdateStatus(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	UpdatePayment(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	FindAll(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) []domain.Order
	FindId(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error)
	DeleteTable(ctx context.Context, tx *sql.Tx)
}
