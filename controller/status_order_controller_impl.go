package controller

import (
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type StatusOrderControllerImpl struct {
	Service service.StatusOrderService
}

func NewStatusOrderController(service service.StatusOrderService) StatusOrderController {
	return &StatusOrderControllerImpl{
		Service: service,
	}
}

func (controller *StatusOrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	statusOrderCreate := web.StatusOrderCreate{}
	helper.ReadFromRequestBody(request, &statusOrderCreate)

	statusOrderResponse := controller.Service.Create(request.Context(), statusOrderCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   statusOrderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusOrderControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	statusOrderUpdate := web.StatusOrderUpdate{}
	helper.ReadFromRequestBody(request, &statusOrderUpdate)

	statusOrderId := params.ByName("statusId")

	statusOrderUpdate.IdStatusOrder = statusOrderId

	statusOrderResponse := controller.Service.Update(request.Context(), statusOrderUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   statusOrderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusOrderControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	statusOrderId := params.ByName("statusId")

	controller.Service.Delete(request.Context(), statusOrderId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusOrderControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	statusOrderId := params.ByName("statusId")

	statudOrderResponse := controller.Service.FindById(request.Context(), statusOrderId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   statudOrderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusOrderControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	statudOrderResponses := controller.Service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   statudOrderResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
