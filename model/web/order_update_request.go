package web

type OrderUpdateRequest struct {
	UserId   int    `json:"user_id,omitempty"`
	OrderId  int    `json:"order_id,omitempty"`
	Payment  string `json:"payment,omitempty"`
	StatusId int    `json:"status_id,omitempty"`
	Image    string `json:"image,omitempty"`
}
