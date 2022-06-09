package web

import (
	"time"

	"github.com/faridlan/restful-api-market/model"
)

type OrdersDetail struct {
	ProductName string `json:"product_name,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
	TotalPrice  int    `json:"total_price,omitempty"`
}

type OrderResponse struct {
	OrderId int    `json:"order_id,omitempty"`
	IdOrder string `json:"id_order,omitempty"`
	// User      *User               `json:"user,omitempty"`
	Address   *AddressReponse     `json:"address,omitempty"`
	Products  []OrdersDetail      `json:"products,omitempty"`
	Total     int                 `json:"total,omitempty"`
	OrderDate time.Time           `json:"order_date,omitempty"`
	Status    StatusOrderResponse `json:"status,omitempty"`
	Payment   *model.NullString   `json:"payment"`
}

type OrderResponseImg struct {
	Image string `json:"image,omitempty"`
}
