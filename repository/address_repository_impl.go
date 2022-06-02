package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepository() AddressRepository {
	return AddressRepositoryImpl{}
}

func (repository AddressRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "insert into addresses (user_id, name, handphone_number, street, districk, post_code, comment) values (?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, address.User.Id, address.Name, address.HandphoneNumber, address.Street, address.Districk, address.PostCode, address.Comment)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	address.Id = int(id)
	return address
}

func (repository AddressRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "update addresses set name = ?, handphone_number = ?, street = ?, districk = ?, post_code = ?, comment = ? where id =? and user_id =?"
	_, err := tx.ExecContext(ctx, SQL, address.Name, address.HandphoneNumber, address.Street, address.Districk, address.PostCode, address.Comment, address.Id, address.User.Id)
	helper.PanicIfError(err)

	return address
}

func (repository AddressRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, address domain.Address) {
	SQL := "delete from addresses where id =? and user_id =?"
	_, err := tx.ExecContext(ctx, SQL, address.Id, address.User.Id)
	helper.PanicIfError(err)
}

func (repository AddressRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, addressId int, userId int) (domain.Address, error) {
	SQL := `select a.id, u.id, u.username, a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment from addresses as a
	left join users as u on u.id = a.user_id
	where a.id = ? and a.user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, addressId, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	address := domain.Address{}

	if rows.Next() {
		err := rows.Scan(&address.Id, &address.User.Id, &address.User.Username, &address.Name, &address.HandphoneNumber, &address.Street, &address.Districk, &address.PostCode, &address.Comment)
		helper.PanicIfError(err)

		return address, nil
	} else {
		return address, errors.New("address not found")
	}
}

func (repository AddressRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Address, error) {
	SQL := `select a.id, u.id, u.username, a.name, a.handphone_number, a.street, a.districk, a.post_code, a.comment from addresses as a
	left join users as u on u.id = a.user_id
	where a.user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	addresses := []domain.Address{}

	for rows.Next() {
		address := domain.Address{}
		err = rows.Scan(&address.Id, &address.User.Id, &address.User.Username, &address.Name, &address.HandphoneNumber, &address.Street, &address.Districk, &address.PostCode, &address.Comment)
		if err != nil {
			if err == sql.ErrNoRows {
				return addresses, errors.New("address not found")
			}
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
