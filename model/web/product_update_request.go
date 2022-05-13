package web

type ProductUpateRequest struct {
	Id          int    `json:"id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	CategoryId  int    `json:"category_id,omitempty"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}