package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ShippingAddressController interface {
	CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
