package controller

import (
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func NewAddressController(addressService service.AddressService) AddressController {
	return &AddressControllerImpl{
		AddressService: addressService,
	}
}

func (controller *AddressControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	addressCreateRequest := web.AddressCreateRequest{}
	helper.ReadFromRequestBody(request, &addressCreateRequest)

	addressCreateRequest.UserId = claim.Id

	addressResponse := controller.AddressService.Create(request.Context(), addressCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	addressUpdateRequest := web.AddressUpdateRequest{}
	helper.ReadFromRequestBody(request, &addressUpdateRequest)

	addressId := params.ByName("addressId")

	addressUpdateRequest.IdAddress = addressId
	addressUpdateRequest.UserId = claim.Id
	addressResponse := controller.AddressService.Update(request.Context(), addressUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	addressId := params.ByName("addressId")

	controller.AddressService.Delete(request.Context(), addressId, claim.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	addressId := params.ByName("addressId")

	addressResponse := controller.AddressService.FindById(request.Context(), addressId, claim.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)
	pagination := helper.Pagination(request)

	addressResponses := controller.AddressService.FindAll(request.Context(), claim.Id, pagination)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
