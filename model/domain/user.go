package domain

import "github.com/faridlan/restful-api-market/model"

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	ImageUrl *model.NullString
	RoleId   int
}
