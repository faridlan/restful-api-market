package web

type UserCreateRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username,omitempty"`
	Email    string `validate:"required,max=100,min=1" json:"email,omitempty"`
	Password string `validate:"required,max=225,min=6" json:"password,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	RoleId   int    `json:"role_id,omitempty"`
	IdRole   string `json:"id_role,omitempty"`
}
