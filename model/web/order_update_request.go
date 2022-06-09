package web

type OrderUpdateRequest struct {
	IdOrder  string `json:"id_order,omitempty"`
	UserId   int    `json:"user_id,omitempty"`
	OrderId  int    `json:"order_id,omitempty"`
	Payment  string `json:"payment,omitempty"`
	StatusId int    `json:"status_id,omitempty"`
	IdStatus string `json:"id_status,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
