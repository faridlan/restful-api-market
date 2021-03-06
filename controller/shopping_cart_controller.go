package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ShoppingCartController interface {
	AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateQty(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindSome(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
