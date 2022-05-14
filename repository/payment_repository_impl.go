package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type PaymentRepositoryImpl struct {
}

func (repository PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "insert into payment(image_url) values (?)"
	resull, err := tx.ExecContext(ctx, SQL, payment.ImageUrl)
	helper.PanicIfError(err)

	id, err := resull.LastInsertId()
	helper.PanicIfError(err)

	payment.Id = int(id)

	return payment
}
