package domain

type Product struct {
	Id          int
	IdProduct   string
	ProductName string
	Category    Category
	Price       int
	Quantity    int
	ImageUrl    string
}
