package web

type AddressCreateRequest struct {
	UserId          int    `json:"user_id,omitempty"`
	Name            string `json:"name,omitempty"`
	HandphoneNumber string `json:"handphone_number,omitempty"`
	Street          string `json:"street,omitempty"`
	Districk        string `json:"districk,omitempty"`
	PostCode        int    `json:"post_code,omitempty"`
	Comment         string `json:"comment,omitempty"`
}
