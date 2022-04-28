package domain

type Address struct {
	Id              int
	User            User
	Name            string
	HandphoneNumber string
	Street          string
	Districk        string
	PostCode        int
	Comment         string
}
