package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type CartRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Update(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Delete(ctx context.Context, tx *sql.Tx, userId int, productId int)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Cart, error)
	FindAll(ctx context.Context, tx *sql.Tx, userId int) []domain.Cart
}
