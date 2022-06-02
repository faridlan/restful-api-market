package web

type ProductResponse struct {
	Id          int               `json:"id,omitempty"`
	IdProduct   string            `json:"id_product,omitempty"`
	ProductName string            `json:"product_name,omitempty"`
	Category    *CategoryResponse `json:"category,omitempty"`
	Price       int               `json:"price,omitempty"`
	Quantity    int               `json:"quantity,omitempty"`
	ImageUrl    string            `json:"image_url,omitempty"`
}
