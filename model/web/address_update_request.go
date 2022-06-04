package web

type AddressUpdateRequest struct {
	Id              int    `json:"id,omitempty"`
	IdAddress       string `json:"id_address,omitempty"`
	UserId          int    `json:"user_id,omitempty"`
	Name            string `json:"name,omitempty" validate:"required"`
	HandphoneNumber string `json:"handphone_number,omitempty" validate:"required"`
	Street          string `json:"street,omitempty" validate:"required"`
	Districk        string `json:"districk,omitempty" validate:"required"`
	PostCode        int    `json:"post_code,omitempty" validate:"required"`
	Comment         string `json:"comment,omitempty" validate:"required"`
}
