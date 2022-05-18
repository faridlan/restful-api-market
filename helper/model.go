package helper

import (
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
		RoleId:   user.RoleId,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToJwtResponse(user web.Claims) web.Claims {
	return web.Claims{
		Id:               user.Id,
		Username:         user.Username,
		Email:            user.Email,
		RoleId:           user.RoleId,
		Token:            user.Token,
		RegisteredClaims: user.RegisteredClaims,
	}
}

func ToBlacklistResponse(blaclist domain.Blacklist) web.BlacklistResponse {
	return web.BlacklistResponse{
		Id:    blaclist.Id,
		Token: blaclist.Token,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		ProductName: product.ProductName,
		Category: web.CategoryResponse{
			Id:           product.Category.Id,
			CategoryName: product.Category.CategoryName,
		},
		Price:    product.Price,
		Quantity: product.Quantity,
		ImageUrl: product.ImageUrl,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		Id: cart.Id,
		User: web.UserResponse{
			Id:       cart.User.Id,
			Username: cart.User.Username,
		},
		Product: web.ProductResponse{
			Id:          cart.Product.Id,
			ProductName: cart.Product.ProductName,
			Category: web.CategoryResponse{
				Id:           cart.Product.Category.Id,
				CategoryName: cart.Product.Category.CategoryName,
			},
			Price:    cart.Product.Price,
			Quantity: cart.Product.Quantity,
			ImageUrl: cart.Product.ImageUrl,
		},
		Quantity: cart.Quantity,
	}
}

func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartResponses []web.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, ToCartResponse(cart))
	}

	return cartResponses
}

func ToAddressResponse(address domain.Address) web.AddressReponse {
	return web.AddressReponse{
		Id: address.Id,
		User: web.UserResponse{
			Id:       address.User.Id,
			Username: address.User.Username,
		},
		Name:            address.Name,
		HandphoneNumber: address.HandphoneNumber,
		Street:          address.Street,
		Districk:        address.Districk,
		PostCode:        address.PostCode,
		Comment:         address.Comment,
	}
}

func ToAddressResponses(addresses []domain.Address) []web.AddressReponse {
	var addressResponses []web.AddressReponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, ToAddressResponse(address))
	}

	return addressResponses
}

func ToOrderDetailResponse(order domain.OrderDetail) web.OrdersDetail {
	return web.OrdersDetail{
		ProductName: order.Product.ProductName,
		Quantity:    order.Quantity,
		Price:       order.UnitPrice,
		TotalPrice:  order.TotalPrice,
	}
}

func ToOrderDetailResponses(orders []domain.OrderDetail) []web.OrdersDetail {
	var orderDetailResponses []web.OrdersDetail
	for _, order := range orders {
		orderDetailResponses = append(orderDetailResponses, ToOrderDetailResponse(order))
	}

	return orderDetailResponses
}

func ToOrderResponse(order domain.Order, orders []web.OrdersDetail) web.OrderResponse {
	return web.OrderResponse{
		OrderId: order.Id,
		User: web.UserResponse{
			Id:       order.User.Id,
			Username: order.User.Username,
		},
		Address: web.AddressReponse{
			// User:            web.UserResponse{},
			Id:              order.Address.Id,
			Name:            order.Address.Name,
			HandphoneNumber: order.Address.HandphoneNumber,
			Street:          order.Address.Street,
			Districk:        order.Address.Districk,
			PostCode:        order.Address.PostCode,
			Comment:         order.Address.Comment,
		},
		Detail:    orders,
		Total:     order.Total,
		OrderDate: order.OrderDate,
		Status: web.StatusOrderResponse{
			Id:         order.Status.Id,
			StatusName: order.Status.StatusName,
		},
		Payment: web.PaymentResponse{
			Id:       order.Payment.Id,
			ImageUrl: order.Payment.ImageUrl,
		},
	}
}

func ToOrdersResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		OrderId: order.Id,
		User: web.UserResponse{
			Id:       order.User.Id,
			Username: order.User.Username,
		},
		Total:     order.Total,
		OrderDate: order.OrderDate,
		Status: web.StatusOrderResponse{
			Id:         order.Status.Id,
			StatusName: order.Status.StatusName,
		},
		// Payment:   web.PaymentResponse{},
	}
}

func ToOrdersResponses(orders []domain.Order) []web.OrderResponse {
	ordersResponses := []web.OrderResponse{}
	for _, order := range orders {
		ordersResponses = append(ordersResponses, ToOrdersResponse(order))
	}

	return ordersResponses
}

func ToCreateOrder(order web.CreateOrder) domain.OrderDetail {
	return domain.OrderDetail{
		Order: domain.Order{
			User: domain.User{
				Id: order.UserId,
			},
		},
		Product: domain.Product{
			Id: order.ProductId,
		},
		Quantity: order.Quantity,
	}
}

func ToCreateOrders(orders []web.CreateOrder) []domain.OrderDetail {
	createOrders := []domain.OrderDetail{}
	for _, order := range orders {
		createOrders = append(createOrders, ToCreateOrder(order))
	}

	return createOrders
}

func ToCartDelete(cart web.CartDeleteRequest) domain.Cart {
	return domain.Cart{
		User: domain.User{
			Id: cart.UserId,
		},
		Product: domain.Product{
			Id: cart.ProductId,
		},
	}
}

func ToCartsDelete(carts []web.CartDeleteRequest) []domain.Cart {
	cartRequest := []domain.Cart{}
	for _, cart := range carts {
		cartRequest = append(cartRequest, ToCartDelete(cart))
	}

	return cartRequest
}

func ToDeleteOrderCart(cart web.CreateOrder) domain.Cart {
	return domain.Cart{
		User: domain.User{
			Id: cart.UserId,
		},
		Product: domain.Product{
			Id: cart.ProductId,
		},
	}
}

func ToDeleteOrderCarts(carts []web.CreateOrder) []domain.Cart {
	cartRequest := []domain.Cart{}
	for _, cart := range carts {
		cartRequest = append(cartRequest, ToDeleteOrderCart(cart))
	}

	return cartRequest
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:           category.Id,
		CategoryName: category.CategoryName,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	categoryResponses := []web.CategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToStatusOrderResponse(statusOrder domain.StatusOrder) web.StatusOrderResponse {
	return web.StatusOrderResponse{
		Id:         statusOrder.Id,
		StatusName: statusOrder.StatusName,
	}
}

func ToStatusOrderResponses(statusOrders []domain.StatusOrder) []web.StatusOrderResponse {
	statusOrderResponses := []web.StatusOrderResponse{}
	for _, statusOrder := range statusOrders {
		statusOrderResponses = append(statusOrderResponses, ToStatusOrderResponse(statusOrder))
	}
	return statusOrderResponses

}
