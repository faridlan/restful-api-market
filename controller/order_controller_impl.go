package controller

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type ShippingAddressControllerImpl struct {
	ShippingAddressService service.ShippingAddressService
}

func NewShippingAddressController(service service.ShippingAddressService) ShippingAddressController {
	return &ShippingAddressControllerImpl{
		ShippingAddressService: service,
	}
}

func (controller *ShippingAddressControllerImpl) CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	detailOrders := []web.CreateOrder{}

	for _, order := range orderCreateRequest.Detail {
		order.UserId = claim.Id
		detailOrders = append(detailOrders, order)
	}

	orderCreate := web.OrderCreateRequest{
		UserId:    claim.Id,
		AddressId: orderCreateRequest.AddressId,
		Detail:    detailOrders,
	}

	orderResponse := controller.ShippingAddressService.CreateOrder(request.Context(), orderCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShippingAddressControllerImpl) FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.ShippingAddressService.FindOrderById(request.Context(), id, claim.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShippingAddressControllerImpl) FindAllOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	orderResponses := controller.ShippingAddressService.FindAllOrderByUser(request.Context(), claim.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShippingAddressControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderUpdateRequest := web.OrderUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderUpdateRequest)

	orderResponse := controller.ShippingAddressService.UpdateStatus(request.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShippingAddressControllerImpl) UpdatePayment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	orderUpdateRequest := web.OrderUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderUpdateRequest)

	orderUpdateRequest.StatusId = 2
	orderUpdateRequest.UserId = claim.Id

	orderResponse := controller.ShippingAddressService.UpdatePayment(request.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShippingAddressControllerImpl) CreateImg(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	file, _, err := request.FormFile("paymentImage")
	helper.PanicIfError(err)
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	helper.PanicIfError(err)

	image := web.OrderUpdateRequest{
		Image: string(fileBytes),
	}

	productResponse := controller.ShippingAddressService.UploadImage(request.Context(), image)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
