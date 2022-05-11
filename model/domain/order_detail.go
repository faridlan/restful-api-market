package domain

type OrderDetail struct {
	Order      Order   `json:"order,omitempty"`
	Product    Product `json:"product,omitempty"`
	Quantity   int     `json:"quantity,omitempty"`
	UnitPrice  int     `json:"unit_price,omitempty"`
	TotalPrice int     `json:"total_price,omitempty"`
}
