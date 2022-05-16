package web

type CreateOrder struct {
	UserId    int `json:"user_id,omitempty" validator:"required"`
	CartId    int `json:"cart_id,omitempty" validator:"required, numeric"`
	ProductId int `json:"product_id,omitempty" validator:"required, numeric"`
	Quantity  int `json:"quantity,omitempty" validator:"required, numeric"`
}

type OrderCreateRequest struct {
	UserId    int `json:"user_id,omitempty"`
	AddressId int `json:"address_id,omitempty" validator:"required, numeric"`
	// ProductId int `json:"product_id,omitempty"`
	// Quantity  int `json:"quantity,omitempty"`
	Detail []CreateOrder `json:"detail,omitempty"`
}
