package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
