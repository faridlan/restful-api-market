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
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	//Admin
	//Product Feature
	productRepository := repository.NewProdcutRepository()
	productService := service.NewProductServie(productRepository, db)
	productController := controller.NewProductController(productService)

	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db)
	CategoryController := controller.NewCategoryController(categoryService)

	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.GET("/api/categories", CategoryController.FindAll)

	//auth
	userRepository := repository.NewUserRepository()
	blacklistRepository := repository.NewBlacklistRepository()
	userService := service.NewAuthService(userRepository, blacklistRepository, db, validate)
	userController := controller.NewAuthController(userService)

	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Register)
	router.POST("/api/logout", userController.Logout)
	router.GET("/api/profile", userController.Profile)
	router.PUT("/api/profile/:userId", userController.UpdateProfile)

	//home
	homeService := service.NewHomeService(productRepository, db)
	homeController := controller.NewHomeController(homeService)

	router.GET("/api/products", homeController.Product)

	//product detail
	cartRepository := repository.NewCartRepository()
	productDetailService := service.NewProductDetailService(productRepository, cartRepository, db)
	productDetailController := controller.NewProductDetailController(productDetailService)

	router.GET("/api/products/:productId", productDetailController.FindProduct)
	router.POST("/api/carts", productDetailController.AddToCart)

	//shopping cart
	shoppingCartService := service.NewShoppingCartService(cartRepository, db)
	shoppingCartController := controller.NewShoppingCartController(shoppingCartService)

	router.GET("/api/carts", shoppingCartController.FindCart)
	router.PUT("/api/carts/:productId", shoppingCartController.UpdateQty)
	router.DELETE("/api/carts", shoppingCartController.DeleteCart)

	//address
	addressRepository := repository.NewAddressRepository()
	addressSerice := service.NewAddressService(addressRepository, db)
	addressController := controller.NewAddressController(addressSerice)

	router.POST("/api/addresses", addressController.Create)
	router.GET("/api/addresses/:addressId", addressController.FindById)
	router.GET("/api/addresses", addressController.FindAll)
	router.PUT("/api/addresses/:addressId", addressController.Update)
	router.DELETE("/api/addresses/:addressId", addressController.Delete)

	//shipping address

	orderRepository := repository.NewOrderRepository()
	orderDetailRepository := repository.NewOrderDetailRepository()
	shippingAddressService := service.NewShippingAddressService(productRepository, orderRepository, orderDetailRepository, cartRepository, db)
	shippingAddressController := controller.NewShippingAddressController(shippingAddressService)

	router.POST("/api/orders", shippingAddressController.CreateOrder)
	router.GET("/api/orders/:orderId", shippingAddressController.FindOrderById)

	// router.PanicHandler = exception.ExceptionError

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
