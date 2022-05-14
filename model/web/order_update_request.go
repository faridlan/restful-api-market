package web

type OrderUpdateRequest struct {
	OrderId   int `json:"order_id,omitempty"`
	PaymentId int `json:"payment_id,omitempty"`
}
