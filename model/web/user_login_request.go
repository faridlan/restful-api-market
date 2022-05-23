package web

type LoginCreateRequest struct {
	Username string ` json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `validate:"required,max=225,min=4" json:"password,omitempty"`
}
