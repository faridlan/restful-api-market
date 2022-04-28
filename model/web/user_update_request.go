package web

type UserUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `validate:"required,max=100,min=1" json:"username,omitempty"`
	Email    string `validate:"required" json:"email,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
