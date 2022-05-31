package model

type Endpoint struct {
	Url    string
	Method string
}

func Endpoints() []Endpoint {
	url := []Endpoint{
		{
			Url:    "/api/categories",
			Method: "POST",
		},
		{
			Url:    "/api/categories/:categoryId",
			Method: "PUT",
		},
		{
			Url:    "/api/categories/:categoryId",
			Method: "GET",
		},
		{
			Url:    "/api/categories",
			Method: "GET",
		},
		{
			Url:    "/api/products",
			Method: "POST",
		},
		{
			Url:    "/api/products/image",
			Method: "POST",
		},
		{
			Url:    "/api/products/:productId",
			Method: "PUT",
		},
		{
			Url:    "/api/statusOrder",
			Method: "POST",
		},
		{
			Url:    "/api/statusOrder/:statusId",
			Method: "PUT",
		},
		{
			Url:    "/api/statusOrder",
			Method: "GET",
		},
		{
			Url:    "/api/statusOrder/:statudId",
			Method: "GET",
		},
		{
			Url:    "/api/roles",
			Method: "POST",
		},
		{
			Url:    "/api/roles/:roleId",
			Method: "PUT",
		},
		{
			Url:    "/api/roles",
			Method: "GET",
		},
	}

	return url
}
