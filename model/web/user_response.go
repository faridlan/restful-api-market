package web

import "github.com/faridlan/restful-api-market/model"

type UserResponse struct {
	Id       int               `json:"id,omitempty"`
	IdUser   string            `json:"id_user,omitempty"`
	Username string            `json:"username,omitempty"`
	Email    string            `json:"email,omitempty"`
	ImageUrl *model.NullString `json:"image_url,omitempty"`
	Role     *RoleResponse     `json:"role,omitempty"`
	Token    string            `json:"token,omitempty"`
}

type UserResponseImg struct {
	ImageUrl string `json:"image_url,omitempty"`
}
