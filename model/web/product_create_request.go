package web

type ProductCreateRequest struct {
	ProductName string `json:"product_name,omitempty" validate:"required"`
	CategoryId  int    `json:"category_id,omitempty" validate:"required"`
	Price       int    `json:"price,omitempty" validate:"required"`
	Quantity    int    `json:"quantity,omitempty" validate:"required"`
	ImageUrl    string `json:"image_url,omitempty" validate:"required"`
}
