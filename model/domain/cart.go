package domain

type Cart struct {
	Id       int
	User     User
	Product  Product
	Quantity int
}
