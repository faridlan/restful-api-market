package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type CartRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Update(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Delete(ctx context.Context, tx *sql.Tx, carts []domain.Cart)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Cart, error)
	FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Cart, error)
	FindSome(ctx context.Context, tx *sql.Tx, cartId []domain.Cart) []domain.Cart
}
