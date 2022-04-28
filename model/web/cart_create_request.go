package web

type CartCreateRequest struct {
	UserId    int `json:"user_id,omitempty"`
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}
