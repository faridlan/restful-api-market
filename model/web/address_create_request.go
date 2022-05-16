package web

type AddressCreateRequest struct {
	UserId          int    `json:"user_id,omitempty" validator:"required, numeric"`
	Name            string `json:"name,omitempty" validator:"required, max=100, min=0"`
	HandphoneNumber string `json:"handphone_number,omitempty" validator:"required, max=100, min=9"`
	Street          string `json:"street,omitempty" validator:"required, max=100, min=0"`
	Districk        string `json:"districk,omitempty" validator:"required, max=100, min=0"`
	PostCode        int    `json:"post_code,omitempty" validator:"required, numeric"`
	Comment         string `json:"comment,omitempty"`
}
