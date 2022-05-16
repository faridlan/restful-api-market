package main

import (
	"net/http"

	"github.com/faridlan/restful-api-market/app"
	"github.com/faridlan/restful-api-market/controller"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/middleware"
	"github.com/faridlan/restful-api-market/repository"
	"github.com/faridlan/restful-api-market/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	//Admin
	//Product Feature
	productRepository := repository.NewProdcutRepository()
	productService := service.NewProductServie(productRepository, db)
	productController := controller.NewProductController(productService)

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db)
	CategoryController := controller.NewCategoryController(categoryService)

	//auth
	userRepository := repository.NewUserRepository()
	blacklistRepository := repository.NewBlacklistRepository()
	userService := service.NewAuthService(userRepository, blacklistRepository, db, validate)
	authController := controller.NewAuthController(userService)

	//product detail
	cartRepository := repository.NewCartRepository()

	//shopping cart
	shoppingCartService := service.NewShoppingCartService(cartRepository, db)
	shoppingCartController := controller.NewShoppingCartController(shoppingCartService)

	//address
	addressRepository := repository.NewAddressRepository()
	addressSerice := service.NewAddressService(addressRepository, db)
	addressController := controller.NewAddressController(addressSerice)

	//shipping address
	orderRepository := repository.NewOrderRepository()
	orderDetailRepository := repository.NewOrderDetailRepository()
	shippingAddressService := service.NewShippingAddressService(productRepository, orderRepository, orderDetailRepository, cartRepository, db)
	shippingAddressController := controller.NewShippingAddressController(shippingAddressService)

	controller := app.ControllerRouter{
		AddressController:         addressController,
		AuthController:            authController,
		CategoryController:        CategoryController,
		ProductController:         productController,
		ShippingAddressController: shippingAddressController,
		ShoppingCartController:    shoppingCartController,
	}

	blacklist := repository.NewBlacklistRepository()

	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router, blacklist, db),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
