package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type UuidRepositoryImpl struct {
}

func NewUuidRepository() UuidRepository {
	return UuidRepositoryImpl{}
}

func (repository UuidRepositoryImpl) CreteUui(ctx context.Context, tx *sql.Tx) (domain.Uuid, error) {
	SQL := "select REPLACE(UUID(),'-','')"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	uuid := domain.Uuid{}
	if rows.Next() {
		err := rows.Scan(&uuid.Uuid)
		helper.PanicIfError(err)
		return uuid, nil
	} else {
		return uuid, errors.New("failed create uuid")
	}
}
