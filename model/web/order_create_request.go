package web

type CreateOrder struct {
	UserId    int `json:"user_id,omitempty"`
	CartId    int `json:"cart_id,omitempty"`
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}

type OrderCreateRequest struct {
	UserId    int `json:"user_id,omitempty"`
	AddressId int `json:"address_id,omitempty"`
	// ProductId int `json:"product_id,omitempty"`
	// Quantity  int `json:"quantity,omitempty"`
	Detail []CreateOrder `json:"detail,omitempty"`
}
