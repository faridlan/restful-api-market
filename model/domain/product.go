package domain

type Product struct {
	Id          int
	ProductName string
	CategoryId  int
	Category    string
	Price       int
	Quantity    int
	ImageUrl    string
}
