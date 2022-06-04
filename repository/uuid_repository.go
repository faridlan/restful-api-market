package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type UuidRepository interface {
	CreteUui(ctx context.Context, tx *sql.Tx) (domain.Uuid, error)
}
