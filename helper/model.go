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
		CategoryId:  product.CategoryId,
		Category:    product.Category,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageUrl:    product.ImageUrl,
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
			Category:    cart.Product.Category,
			Price:       cart.Product.Price,
			Quantity:    cart.Product.Quantity,
			ImageUrl:    cart.Product.ImageUrl,
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
