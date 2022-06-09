package web

type AddressCreateRequest struct {
	UserId          int    `json:"id_user,omitempty"`
	Name            string `json:"name,omitempty" validate:"required,max=100,min=0"`
	HandphoneNumber string `json:"handphone_number,omitempty" validate:"required,max=100,min=9"`
	Street          string `json:"street,omitempty" validate:"required,max=100,min=0"`
	Districk        string `json:"district,omitempty" validate:"required,max=100,min=0"`
	PostCode        int    `json:"postal_code,omitempty" validate:"required,numeric"`
	Comment         string `json:"comment,omitempty" validate:"required"`
}
