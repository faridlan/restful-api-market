package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type StatusOrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, statusCode domain.StatusCode) domain.StatusCode
	Update(ctx context.Context, tx *sql.Tx, statusCode domain.StatusCode) domain.StatusCode
	Delete(ctx context.Context, tx *sql.Tx, statusCode domain.StatusCode)
	FindById(ctx context.Context, tx *sql.Tx, statusId int) (domain.StatusCode, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.StatusCode
}
