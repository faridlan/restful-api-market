package web

type CreateOrder struct {
	UserId    int `json:"user_id,"`
	CartId    int `json:"cart_id," validate:"required,min=0"`
	ProductId int `json:"product_id," validate:"required,min=0"`
	Quantity  int `json:"quantity," validate:"required,min=0"`
}

type OrderCreateRequest struct {
	UserId    int `json:"user_id" validate:"required"`
	AddressId int `json:"address_id" validate:"required"`
	// ProductId int `json:"product_id,omitempty"`
	// Quantity  int `json:"quantity,omitempty"`
	Detail []CreateOrder `json:"detail"`
}
