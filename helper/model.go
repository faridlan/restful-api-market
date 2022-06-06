package helper

import (
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		IdUser:   user.IdUser,
		Username: user.Username,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
		Role: &web.RoleResponse{
			Id:     user.Role.Id,
			IdRole: user.Role.IdRole,
			Name:   user.Role.Name,
		},
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
		IdUser:           user.IdUser,
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
		IdProduct:   product.IdProduct,
		ProductName: product.ProductName,
		Category: &web.CategoryResponse{
			Id:           product.Category.Id,
			IdCategory:   product.Category.IdCategory,
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
			Category: &web.CategoryResponse{
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
		Id:        address.Id,
		IdAddress: address.IdAddress,
		User: &web.UserResponse{
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
		IdOrder: order.IdOrder,
		User: &web.UserResponse{
			Id:       order.User.Id,
			Username: order.User.Username,
		},
		Address: &web.AddressReponse{
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
		Payment: order.Payment,
	}
}

func ToOrdersResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		OrderId: order.Id,
		IdOrder: order.IdOrder,
		User: &web.UserResponse{
			Id:       order.User.Id,
			Username: order.User.Username,
		},
		Total:     order.Total,
		OrderDate: order.OrderDate,
		Status: web.StatusOrderResponse{
			Id:         order.Status.Id,
			StatusName: order.Status.StatusName,
		},
		Payment: order.Payment,
	}
}

func ToOrdersResponses(orders []domain.Order) []web.OrderResponse {
	ordersResponses := []web.OrderResponse{}
	for _, order := range orders {
		ordersResponses = append(ordersResponses, ToOrdersResponse(order))
	}

	return ordersResponses
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
			IdProduct: cart.ProductId,
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
		IdCategory:   category.IdCategory,
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
		IdStatus:   statusOrder.IdStatusOrder,
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

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:     role.Id,
		IdRole: role.IdRole,
		Name:   role.Name,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	roleResponses := []web.RoleResponse{}
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}

	return roleResponses
}

func ToSelectCartRequest(cart web.CartSelectRequest) domain.Cart {
	return domain.Cart{
		Id: cart.CartId,
	}
}

func ToSelectCartsRequest(carts []web.CartSelectRequest) []domain.Cart {
	cartsRequest := []domain.Cart{}
	for _, c := range carts {
		cartsRequest = append(cartsRequest, ToSelectCartRequest(c))
	}

	return cartsRequest
}

func ToFindProduct(order web.CreateOrder) domain.Product {
	return domain.Product{
		IdProduct: order.ProductId,
	}
}

func ToFindProducts(orders []web.CreateOrder) []domain.Product {
	ordersCreate := []domain.Product{}
	for _, order := range orders {
		ordersCreate = append(ordersCreate, ToFindProduct(order))
	}

	return ordersCreate
}

func ToCreateOrder(order web.CreateOrder) domain.OrderDetail {
	return domain.OrderDetail{
		Order: domain.Order{
			User: domain.User{
				Id: order.UserId,
			},
		},
		Product: domain.Product{
			IdProduct: order.ProductId,
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

// func ToRegisterResponse()  {
// 	return web.Claims{
// 		Id:               user.Id,
// 		IdUser:           user.IdUser,
// 		Username:         user.Username,
// 		Email:            user.Email,
// 		RoleId:           user.RoleId,
// 		Token:            user.Token,
// 		RegisteredClaims: user.RegisteredClaims,
// 	}
// 	return web.UserResponse{
// 		Id:       user.Id,
// 		IdUser:   user.IdUser,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		ImageUrl: user.ImageUrl,
// 		Role: &web.RoleResponse{
// 			Id:     user.Role.Id,
// 			IdRole: user.Role.IdRole,
// 			Name:   user.Role.Name,
// 		},
// 	}
// }
