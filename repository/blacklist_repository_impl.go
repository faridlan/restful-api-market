package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type BlacklistRepositoryImpl struct {
}

func NewBlacklistRepository() BlacklistRepository {
	return BlacklistRepositoryImpl{}
}

func (repository BlacklistRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, blacklist domain.Blacklist) domain.Blacklist {
	SQL := "insert into blacklist (token) values (?)"
	result, err := tx.ExecContext(ctx, SQL, blacklist.Token)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	blacklist.Id = int(id)

	return blacklist
}

func (repository BlacklistRepositoryImpl) SelectById(ctx context.Context, tx *sql.Tx, token string) (domain.Blacklist, error) {
	SQL := "select id,token from blacklist where token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	helper.PanicIfError(err)
	defer rows.Close()

	blacklist := domain.Blacklist{}
	if rows.Next() {
		err := rows.Scan(&blacklist.Id, &blacklist.Token)
		helper.PanicIfError(err)
		return blacklist, nil
	} else {
		return blacklist, errors.New("unauthorized")
	}
}

func (repository BlacklistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, blacklistId int) {
	SQL := "delete from blacklist where id = ?"
	_, err := tx.ExecContext(ctx, SQL, blacklistId)
	helper.PanicIfError(err)
}
