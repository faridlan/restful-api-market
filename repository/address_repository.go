package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type AddressRepository interface {
	Save(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address
	Update(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address
	Delete(ctx context.Context, tx *sql.Tx, address domain.Address)
	FindById(ctx context.Context, tx *sql.Tx, addressId string, userId int) (domain.Address, error)
	FindAll(ctx context.Context, tx *sql.Tx, userId int, pagination domain.Pagination) ([]domain.Address, error)
	FindSeeder(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) (domain.Address, error)
	DeleteTable(ctx context.Context, tx *sql.Tx)
}
