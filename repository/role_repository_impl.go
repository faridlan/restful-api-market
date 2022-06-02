package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return RoleRepositoryImpl{}
}

func (repository RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "insert into roles(id_role, role_name) values (REPLACE(UUID(),'-',''),?)"
	result, err := tx.ExecContext(ctx, SQL, role.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)

	return role
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update roles set role_name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Name, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error) {
	SQL := "select id, id_role, role_name from roles where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)

	defer rows.Close()

	role := domain.Role{}

	if rows.Next() {
		err := rows.Scan(&role.Id, &role.IdRole, &role.Name)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("role not found")
	}

}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, id_role, role_name from roles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	roles := []domain.Role{}

	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.IdRole, &role.Name)
		helper.PanicIfError(err)

		roles = append(roles, role)
	}

	return roles
}
