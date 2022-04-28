package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController interface {
	Product(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
