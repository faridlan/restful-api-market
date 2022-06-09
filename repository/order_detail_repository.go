package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type OrderDetailRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orders []domain.OrderDetail) []domain.OrderDetail
	UpdateTotal(ctx context.Context, tx *sql.Tx, orders []domain.OrderDetail) []domain.OrderDetail
	UpdateProductQty(ctx context.Context, tx *sql.Tx, products []domain.OrderDetail) []domain.OrderDetail
	FindById(ctx context.Context, tx *sql.Tx, orderId int, userId int) []domain.OrderDetail
	AdminFindById(ctx context.Context, tx *sql.Tx, orderId int) []domain.OrderDetail
	DeleteTable(ctx context.Context, tx *sql.Tx)
}
