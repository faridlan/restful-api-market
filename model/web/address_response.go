package web

type AddressReponse struct {
	Id              int           `json:"id,omitempty"`
	IdAddress       string        `json:"id_address,omitempty"`
	User            *UserResponse `json:"user,omitempty"`
	Name            string        `json:"name,omitempty"`
	HandphoneNumber string        `json:"handphone_number,omitempty"`
	Street          string        `json:"street,omitempty"`
	Districk        string        `json:"districk,omitempty"`
	PostCode        int           `json:"post_code,omitempty"`
	Comment         string        `json:"comment,omitempty"`
}
