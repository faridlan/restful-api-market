package app

import (
	"github.com/faridlan/restful-api-market/controller"
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
}

func NewRouter(controller ControllerRouter) *httprouter.Router {

	router := httprouter.New()

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
	router.POST("/api/profiles/image", controller.AuthController.CreateImg) //UPLOAD IMAGE
	router.PUT("/api/profiles/:userId", controller.AuthController.UpdateProfile)
	router.GET("/api/profiles/:userId", controller.AuthController.Profile) //FOR ADMIN
	router.GET("/api/profiles", controller.AuthController.FindAll)         //GET ALL USERS
	router.GET("/api/myprofile", controller.AuthController.MyProfile)

	//CATEGORIES
	router.POST("/api/categories", controller.CategoryController.Create)
	router.PUT("/api/categories/:categoryId", controller.CategoryController.Update)
	// router.DELETE("/api/categories/:categoryId", controller.CategoryController.Delete) //NOT RECOMENDED
	router.GET("/api/categories/:categoryId", controller.CategoryController.FindById)
	router.GET("/api/categories", controller.CategoryController.FindAll)

	//PRODUCTS
	router.GET("/api/products", controller.ProductController.FindById)
	router.POST("/api/products", controller.ProductController.Create)
	router.POST("/api/products/image", controller.ProductController.CreateImg)
	router.PUT("/api/products/:productId", controller.ProductController.Update)
	// router.DELETE("/api/products/:productId", controller.ProductController.Delete) //NOT RECOMENDED
	router.GET("/api/products/:productId", controller.ProductController.FindById)

	//ORDERS
	router.POST("/api/orders", controller.ShippingAddressController.CreateOrder)
	router.GET("/api/orders/:orderId", controller.ShippingAddressController.FindOrderById)
	router.GET("/api/orders", controller.ShippingAddressController.FindAllOrder)
	router.PUT("/api/status/order", controller.ShippingAddressController.UpdateStatus)
	router.PUT("/api/payment/order", controller.ShippingAddressController.UpdatePayment)
	router.POST("/api/payment/image", controller.ShippingAddressController.CreateImg)

	//CARTS
	router.POST("/api/carts", controller.ShoppingCartController.AddToCart)
	router.GET("/api/carts", controller.ShoppingCartController.FindCart)
	// router.GET("/api/carts", controller.ShoppingCartController.FindCart) //FINDSOME CART //UNFINISH
	router.PUT("/api/carts/:productId", controller.ShoppingCartController.UpdateQty)
	router.DELETE("/api/carts", controller.ShoppingCartController.DeleteCart)

	//STATUSORDERS
	router.POST("/api/status-order", controller.StatusOrderController.Create)
	router.PUT("/api/status-order/:statusId", controller.StatusOrderController.Update)
	// router.DELETE("/api/statusOrder/:statusId", controller.StatusOrderController.Delete) //NOT RECOMENDED
	router.GET("/api/status-order/:statusId", controller.StatusOrderController.FindById)
	router.GET("/api/status-order", controller.StatusOrderController.FindAll)

	//ROLE
	router.POST("/api/roles", controller.RoleController.Create)
	router.PUT("/api/roles/:roleId", controller.RoleController.Update)
	router.GET("/api/roles", controller.RoleController.FindAll)

	// router.PanicHandler = exception.ExceptionError

	return router
}
