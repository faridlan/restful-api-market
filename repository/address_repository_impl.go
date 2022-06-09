package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepository() AddressRepository {
	return AddressRepositoryImpl{}
}

func (repository AddressRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "insert into addresses (id_address, user_id, name, handphone_number, street, districk, post_code, comment) values (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, address.IdAddress, address.User.Id, address.Name, address.HandphoneNumber, address.Street, address.Districk, address.PostCode, address.Comment)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	address.Id = int(id)
	return address
}

func (repository AddressRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "update addresses set name = ?, handphone_number = ?, street = ?, districk = ?, post_code = ?, comment = ? where id_address =? and user_id =?"
	_, err := tx.ExecContext(ctx, SQL, address.Name, address.HandphoneNumber, address.Street, address.Districk, address.PostCode, address.Comment, address.IdAddress, address.User.Id)
	helper.PanicIfError(err)

	return address
}

func (repository AddressRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, address domain.Address) {
	SQL := "delete from addresses where id =? and user_id =?"
	_, err := tx.ExecContext(ctx, SQL, address.Id, address.User.Id)
	helper.PanicIfError(err)
}

func (repository AddressRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, addressId string, userId int) (domain.Address, error) {
	SQL := `select a.id ,a.id_address,u.id, u.username, a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment from addresses as a
	left join users as u on u.id = a.user_id
	where a.id_address = ? and a.user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, addressId, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	address := domain.Address{}

	if rows.Next() {
		err := rows.Scan(&address.Id, &address.IdAddress, &address.User.Id, &address.User.Username, &address.Name, &address.HandphoneNumber, &address.Street, &address.Districk, &address.PostCode, &address.Comment)
		helper.PanicIfError(err)

		return address, nil
	} else {
		return address, errors.New("address not found")
	}
}

func (repository AddressRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int, pagination domain.Pagination) ([]domain.Address, error) {
	SQL := fmt.Sprintf(`select a.id_address, u.id, u.username, a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment from addresses as a
	left join users as u on u.id = a.user_id
	where a.user_id = ?
	order by a.id desc limit %d,%d`, pagination.Page, pagination.Limit)
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	addresses := []domain.Address{}

	for rows.Next() {
		address := domain.Address{}
		err := rows.Scan(&address.IdAddress, &address.User.Id, &address.User.Username, &address.Name, &address.HandphoneNumber, &address.Street, &address.Districk, &address.PostCode, &address.Comment)
		helper.PanicIfError(err)
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (repository AddressRepositoryImpl) FindSeeder(ctx context.Context, tx *sql.Tx, pagination domain.Pagination) (domain.Address, error) {
	SQL := fmt.Sprintf("select id, id_address from addresses order by id limit %d,%d", pagination.Page, pagination.Limit)
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	address := domain.Address{}
	if rows.Next() {
		err := rows.Scan(&address.Id, &address.IdAddress)
		helper.PanicIfError(err)
		return address, nil
	} else {
		return address, errors.New("user not found")
	}
}

func (repository AddressRepositoryImpl) DeleteTable(ctx context.Context, tx *sql.Tx) {
	SQL := "delete from addresses"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
