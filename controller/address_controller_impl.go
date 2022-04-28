package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func (controller *AddressControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	addressCreateRequest := web.AddressCreateRequest{}
	helper.ReadFromRequestBody(request, &addressCreateRequest)

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
	id, err := strconv.Atoi(addressId)
	helper.PanicIfError(err)

	addressUpdateRequest.Id = id
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
	id, err := strconv.Atoi(addressId)
	helper.PanicIfError(err)

	controller.AddressService.Delete(request.Context(), id, claim.Id)

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
	id, err := strconv.Atoi(addressId)
	helper.PanicIfError(err)

	addressResponse := controller.AddressService.FindById(request.Context(), id, claim.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
