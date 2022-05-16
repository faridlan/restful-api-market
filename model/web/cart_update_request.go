package web

type CartUpdateRequest struct {
	UserId    int `json:"user_id" validate:"required,numeric"`
	ProductId int ` json:"product_id"`
	Quantity  int `validate:"required,numeric,min=0" json:"quantity"`
}
