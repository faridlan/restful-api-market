package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type StatusOrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, statusCode domain.StatusOrder) domain.StatusOrder
	Update(ctx context.Context, tx *sql.Tx, statusCode domain.StatusOrder) domain.StatusOrder
	Delete(ctx context.Context, tx *sql.Tx, statusCode domain.StatusOrder)
	FindById(ctx context.Context, tx *sql.Tx, statusId int) (domain.StatusOrder, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.StatusOrder
}
