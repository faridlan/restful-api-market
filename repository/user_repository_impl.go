package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users (username,email,password) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "select id,username,email,role_id from users where (username =? or email =?) and password =?"
	// SQL := "select id,username,email from users where (username =? or email =?) and password =?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	defer rows.Close()
	user = domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Role.Id)
		// err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("unauthorized")
	}
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := `select u.id,u.username,u.email,u.image_url,r.id, r.role_name from users as u 
	inner join roles as r on u.role_id = r.id 
	where u.id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()
	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.ImageUrl, &user.Role.Id, &user.Role.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `select u.id,u.username,u.email,u.image_url,r.id, r.role_name from users as u 
	inner join roles as r on u.role_id = r.id`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.ImageUrl, &user.Role.Id, &user.Role.Name)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set username = ? , email = ?, image_url = ? where id = ?"
	// SQL := "update users set username = ? , email = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.ImageUrl, user.Id)
	helper.PanicIfError(err)

	return user
}
