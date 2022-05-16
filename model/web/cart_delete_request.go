package web

type CartDeleteRequest struct {
	UserId    int `json:"user_id,omitempty"`
	ProductId int `json:"product_id,omitempty" validate:"required"`
}

type CartsDeleteRequest struct {
	Detail []CartDeleteRequest `json:"detail,omitempty"`
}
