package web

type AddressReponse struct {
	Id        int    `json:"id,omitempty"`
	IdAddress string `json:"id_address,omitempty"`
	// User            User   `json:"user,omitempty"`
	Name            string `json:"name,omitempty"`
	HandphoneNumber string `json:"handphone_number,omitempty"`
	Street          string `json:"street,omitempty"`
	Districk        string `json:"district,omitempty"`
	PostCode        int    `json:"postal_code,omitempty"`
	Comment         string `json:"comment,omitempty"`
}
