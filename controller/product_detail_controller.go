package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductDetailController interface {
	FindProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
