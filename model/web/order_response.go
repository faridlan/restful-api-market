package web

import "time"

type OrdersDetail struct {
	ProductName string `json:"product_name,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
	TotalPrice  int    `json:"total_price,omitempty"`
}

type OrderResponse struct {
	OrderId   int            `json:"order_id,omitempty"`
	User      UserResponse   `json:"user,omitempty"`
	Address   AddressReponse `json:"address,omitempty"`
	Detail    []OrdersDetail `json:"detail,omitempty"`
	Total     int            `json:"total,omitempty"`
	OrderDate time.Time      `json:"order_date,omitempty"`
}
