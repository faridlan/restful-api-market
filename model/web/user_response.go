package web

import "github.com/faridlan/restful-api-market/model"

type UserResponse struct {
	Id       int               `json:"id,omitempty"`
	Username string            `json:"username,omitempty"`
	Email    string            `json:"email,omitempty"`
	ImageUrl *model.NullString `json:"image_url"`
	RoleId   int               `json:"role_id,omitempty"`
}
