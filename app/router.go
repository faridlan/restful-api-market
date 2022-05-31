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
}

func NewRouter(controller ControllerRouter) *httprouter.Router {

	router := httprouter.New()

	//ADDRESS
	router.POST("/api/users/addresses", controller.AddressController.Create)
	router.GET("/api/users/addresses/:addressId", controller.AddressController.FindById)
	router.GET("/api/users/addresses", controller.AddressController.FindAll)
	router.PUT("/api/users/addresses/:addressId", controller.AddressController.Update)
	router.DELETE("/api/users/addresses/:addressId", controller.AddressController.Delete)

	//USERS OR AUTH
	router.POST("/api/users/register", controller.AuthController.Register)
	router.POST("/api/users/login", controller.AuthController.Login)
	router.POST("/api/users/logout", controller.AuthController.Logout)
	// router.PUT("/api/users/profile/:userId", controller.AuthController.Logout) //FOR ADMIN
	// router.POST("/api/users/profile/image", controller.AuthController.Logout) //UPLOAD IMAGE
	// router.POST("/api/users/profile/image", controller.AuthController.Logout) //GET ALL USERS
	router.GET("/api/users/profile", controller.AuthController.Profile)

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
	router.POST("/api/statusOrder", controller.StatusOrderController.Create)
	router.PUT("/api/statusOrder/:statusId", controller.StatusOrderController.Update)
	// router.DELETE("/api/statusOrder/:statusId", controller.StatusOrderController.Delete) //NOT RECOMENDED
	router.GET("/api/statusOrder/:statusId", controller.StatusOrderController.FindById)
	router.GET("/api/statusOrder", controller.StatusOrderController.FindAll)

	//ROLE

	// router.PanicHandler = exception.ExceptionError

	return router
}
