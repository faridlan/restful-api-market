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

	uuidRepository := repository.NewUuidRepository()

	//Admin
	//Product Feature
	productRepository := repository.NewProdcutRepository()
	productService := service.NewProductServie(productRepository, uuidRepository, db, validate)
	productController := controller.NewProductController(productService)

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, uuidRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	//auth
	userRepository := repository.NewUserRepository()
	blacklistRepository := repository.NewBlacklistRepository()
	userService := service.NewAuthService(userRepository, blacklistRepository, uuidRepository, db, validate)
	authController := controller.NewAuthController(userService)

	//product detail
	cartRepository := repository.NewCartRepository()

	//shopping cart
	shoppingCartService := service.NewShoppingCartService(cartRepository, db, validate)
	shoppingCartController := controller.NewShoppingCartController(shoppingCartService)

	//address
	addressRepository := repository.NewAddressRepository()
	addressSerice := service.NewAddressService(addressRepository, uuidRepository, db, validate)
	addressController := controller.NewAddressController(addressSerice)

	//shipping address
	orderRepository := repository.NewOrderRepository()
	orderDetailRepository := repository.NewOrderDetailRepository()
	shippingAddressService := service.NewShippingAddressService(productRepository, orderRepository, orderDetailRepository, cartRepository, addressRepository, uuidRepository, db, validate)
	shippingAddressController := controller.NewShippingAddressController(shippingAddressService)

	//Status Order
	statusOrderRepository := repository.NewStatusOrderRepository()
	statusOrderService := service.NewStatusOrderService(statusOrderRepository, uuidRepository, db, validate)
	statusOrderContoller := controller.NewStatusOrderController(statusOrderService)

	//Roles
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, uuidRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	controller := app.ControllerRouter{
		AddressController:         addressController,
		AuthController:            authController,
		CategoryController:        CategoryController,
		ProductController:         productController,
		ShippingAddressController: shippingAddressController,
		ShoppingCartController:    shoppingCartController,
		StatusOrderController:     statusOrderContoller,
		RoleController:            roleController,
	}

	blacklist := repository.NewBlacklistRepository()
	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.NewAuthMiddleware(router, blacklist, db),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
