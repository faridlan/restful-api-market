package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
}
