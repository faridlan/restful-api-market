package web

type CartResponse struct {
	Id       int             `json:"id,omitempty"`
	User     UserResponse    `json:"user,omitempty"`
	Product  ProductResponse `json:"product,omitempty"`
	Quantity int             `json:"quantity,omitempty"`
}
