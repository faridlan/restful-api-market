package web

type CreateOrder struct {
	UserId    int    `json:"user_id,"`
	ProductId string `json:"product_id," validate:"required,min=0"`
	Quantity  int    `json:"quantity," validate:"required,min=0"`
}

type OrderCreateRequest struct {
	UserId    int           `json:"user_id" validate:"required"`
	AddressId string        `json:"address_id" validate:"required"`
	Detail    []CreateOrder `json:"detail"`
}
