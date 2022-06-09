package web

type CreateOrder struct {
	UserId    int    `json:"id_user,"`
	ProductId string `json:"id_product," validate:"required,min=0"`
	Quantity  int    `json:"quantity," validate:"required,min=0"`
}

type OrderCreateRequest struct {
	UserId    int           `json:"id_user"`
	AddressId string        `json:"id_address" validate:"required"`
	Products  []CreateOrder `json:"products"`
}
