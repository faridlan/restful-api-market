package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type BlacklistRepository interface {
	Create(ctx context.Context, tx *sql.Tx, blacklist domain.Blacklist) domain.Blacklist
	SelectById(ctx context.Context, tx *sql.Tx, token string) (domain.Blacklist, error)
	Delete(ctx context.Context, tx *sql.Tx, blacklistId int)
}
