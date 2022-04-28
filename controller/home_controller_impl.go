package controller

import (
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type HomeControllerImpl struct {
	Service service.HomeService
}

func NewHomeController(service service.HomeService) HomeController {
	return &HomeControllerImpl{
		Service: service,
	}
}

func (controller *HomeControllerImpl) Product(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.Service.Product(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
