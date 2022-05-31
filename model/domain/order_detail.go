package domain

type OrderDetail struct {
	Order      Order
	Product    Product
	Quantity   int
	UnitPrice  int
	TotalPrice int
}
