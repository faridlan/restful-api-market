package app

import (
	"github.com/faridlan/restful-api-market/controller"
	"github.com/faridlan/restful-api-market/exception"
	"github.com/julienschmidt/httprouter"
)

type ControllerRouter struct {
	AddressController         controller.AddressController
	AuthController            controller.AuthController
	CategoryController        controller.CategoryController
	ProductController         controller.ProductController
	ShippingAddressController controller.ShippingAddressController
	ShoppingCartController    controller.ShoppingCartController
	StatusOrderController     controller.StatusOrderController
	RoleController            controller.RoleController
	SeederController          controller.SeederController
	HomeController            controller.HomeController
}

func NewRouter(controller ControllerRouter) *httprouter.Router {

	router := httprouter.New()

	router.GET("/", controller.HomeController.Home)

	//ADDRESS
	router.POST("/api/addresses", controller.AddressController.Create)
	router.GET("/api/addresses/:addressId", controller.AddressController.FindById)
	router.GET("/api/addresses", controller.AddressController.FindAll)
	router.PUT("/api/addresses/:addressId", controller.AddressController.Update)
	router.DELETE("/api/addresses/:addressId", controller.AddressController.Delete)

	//USERS OR AUTH
	router.POST("/api/register", controller.AuthController.Register)
	router.POST("/api/login", controller.AuthController.Login)
	router.POST("/api/logout", controller.AuthController.Logout)
	router.POST("/api/profiles/image", controller.AuthController.CreateImg)
	router.PUT("/api/profiles/:userId", controller.AuthController.UpdateProfile)
	router.GET("/api/profiles", controller.AuthController.MyProfile)
	router.POST("/api/users", controller.AuthController.CreateUsers)    //FOR ADMIN
	router.GET("/api/users/:userId", controller.AuthController.Profile) // FOR ADMIN
	router.GET("/api/users", controller.AuthController.FindAll)         // FOR ADMIN

	//CATEGORIES //FOR ADMIN
	router.POST("/api/categories", controller.CategoryController.Create)
	router.PUT("/api/categories/:categoryId", controller.CategoryController.Update)
	// router.DELETE("/api/categories/:categoryId", controller.CategoryController.Delete) //NOT RECOMENDED
	router.GET("/api/categories/:categoryId", controller.CategoryController.FindById)
	router.GET("/api/categories", controller.CategoryController.FindAll)

	//PRODUCTS
	router.POST("/api/products", controller.ProductController.Create)           // FOR ADMIN
	router.PUT("/api/products/:productId", controller.ProductController.Update) // FOR ADMIN
	router.GET("/api/products", controller.ProductController.FindAll)
	router.GET("/api/products/:productId", controller.ProductController.FindById)
	router.POST("/api/products/image", controller.ProductController.CreateImg)
	// router.DELETE("/api/products/:productId", controller.ProductController.Delete) //NOT RECOMENDED

	//ORDERS
	router.POST("/api/orders", controller.ShippingAddressController.CreateOrder)
	router.POST("/api/payment/image", controller.ShippingAddressController.CreateImg)
	router.GET("/api/customer/orders/:orderId", controller.ShippingAddressController.FindOrderById)
	router.GET("/api/customer/orders", controller.ShippingAddressController.FindAllOrder)
	router.PUT("/api/status/orders/:orderId", controller.ShippingAddressController.UpdateStatus)
	router.PUT("/api/payment/orders/:orderId", controller.ShippingAddressController.UpdatePayment)
	router.GET("/api/orders/:orderId", controller.ShippingAddressController.FindById) // FOR ADMIN
	router.GET("/api/orders", controller.ShippingAddressController.FindAll)           // FOR ADMIN

	//CARTS
	router.POST("/api/carts", controller.ShoppingCartController.AddToCart)
	router.GET("/api/carts", controller.ShoppingCartController.FindCart)
	router.GET("/api/some/carts", controller.ShoppingCartController.FindSome)
	router.PUT("/api/carts/:productId", controller.ShoppingCartController.UpdateQty)
	router.DELETE("/api/carts", controller.ShoppingCartController.DeleteCart)

	//STATUSORDERS //FOR ADMIN
	router.POST("/api/status-order", controller.StatusOrderController.Create)
	router.PUT("/api/status-order/:statusId", controller.StatusOrderController.Update)
	// router.DELETE("/api/statusOrder/:statusId", controller.StatusOrderController.Delete) //NOT RECOMENDED
	router.GET("/api/status-order/:statusId", controller.StatusOrderController.FindById)
	router.GET("/api/status-order", controller.StatusOrderController.FindAll)

	//ROLE // FOR ADMIN
	router.POST("/api/roles", controller.RoleController.Create)
	router.PUT("/api/roles/:roleId", controller.RoleController.Update)
	router.GET("/api/roles/:roleId", controller.RoleController.FindById)
	router.GET("/api/roles", controller.RoleController.FindAll)

	//Seeder
	router.POST("/api/seeder", controller.SeederController.Create)
	router.DELETE("/api/seeder", controller.SeederController.Delete)

	router.PanicHandler = exception.ExceptionError

	return router
}
