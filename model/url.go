package model

type Endpoint struct {
	Url    string
	Method string
}

func Endpoints() []Endpoint {
	url := []Endpoint{
		{
			Url:    "/api/profiles",
			Method: "GET",
		},
		{
			Url:    "/api/profile/:userId",
			Method: "GET",
		},
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
			Url:    "/api/status/order",
			Method: "PUT",
		},
		{
			Url:    "/api/orders",
			Method: "GET",
		},
		{
			Url:    "/api/order/:id",
			Method: "GET",
		},
		{
			Url:    "/api/status-order",
			Method: "POST",
		},
		{
			Url:    "/api/status-order/:statusId",
			Method: "PUT",
		},
		{
			Url:    "/api/status-order",
			Method: "GET",
		},
		{
			Url:    "/api/status-order/:statusId",
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
