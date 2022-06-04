package domain

import "github.com/faridlan/restful-api-market/model"

type User struct {
	Id       int
	IdUser   string
	Username string
	Email    string
	Password string
	ImageUrl *model.NullString
	Role     Role
}
