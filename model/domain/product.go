package domain

type Product struct {
	Id          int
	ProductName string
	Category    Category
	Price       int
	Quantity    int
	ImageUrl    string
}
