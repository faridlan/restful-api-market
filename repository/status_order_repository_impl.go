package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type StatusOrderRepositoryImpl struct {
}

func NewStatusOrderRepository() StatusOrderRepository {
	return StatusOrderRepositoryImpl{}
}

func (repository StatusOrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, statusOrder domain.StatusOrder) domain.StatusOrder {

	SQL := "insert into status_order (status_name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, statusOrder.StatusName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	statusOrder.Id = int(id)

	return statusOrder

}

func (repository StatusOrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, statusCode domain.StatusOrder) domain.StatusOrder {

	SQL := "update status_order set status_name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, statusCode.StatusName, statusCode.Id)
	helper.PanicIfError(err)

	return statusCode

}

func (repository StatusOrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, statusCode domain.StatusOrder) {

	SQL := "delete from status_order where id = ?"
	_, err := tx.ExecContext(ctx, SQL, statusCode.Id)
	helper.PanicIfError(err)

}

func (repository StatusOrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, statusId int) (domain.StatusOrder, error) {

	SQL := "select id, status_name from status_order where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, statusId)
	helper.PanicIfError(err)

	defer rows.Close()

	status := domain.StatusOrder{}

	if rows.Next() {
		err := rows.Scan(&status.Id, &status.StatusName)
		helper.PanicIfError(err)

		return status, nil
	} else {
		return status, errors.New("status not found")
	}

}

func (repository StatusOrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.StatusOrder {

	SQL := "select id, status_name from status_order"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	statusAll := []domain.StatusOrder{}

	for rows.Next() {
		status := domain.StatusOrder{}
		err := rows.Scan(&status.Id, &status.StatusName)
		helper.PanicIfError(err)

		statusAll = append(statusAll, status)
	}

	return statusAll
}
