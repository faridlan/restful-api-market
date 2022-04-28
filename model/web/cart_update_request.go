package web

type CartUpdateRequest struct {
	UserId    int `json:"user_id,omitempty"`
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}
