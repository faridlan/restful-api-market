package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type ShoppingCartControllerImpl struct {
	Service service.ShoppingCartService
}

func NewShoppingCartController(service service.ShoppingCartService) ShoppingCartController {
	return &ShoppingCartControllerImpl{
		Service: service,
	}
}

func (controller *ShoppingCartControllerImpl) AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claim := web.Claims{}
	helper.ParseJwt(request, &claim)
	cartCreateRequest := web.CartCreateRequest{}
	cartCreateRequest.UserId = claim.Id
	helper.ReadFromRequestBody(request, &cartCreateRequest)

	cartResponse := controller.Service.AddToCart(request.Context(), cartCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) FindCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	cartResponses := controller.Service.FindCart(request.Context(), claim.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) UpdateQty(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	cartUpdateRequest := web.CartUpdateRequest{}
	helper.ReadFromRequestBody(request, &cartUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	cartUpdateRequest.UserId = claim.Id
	cartUpdateRequest.ProductId = id

	cartResponse := controller.Service.UpdateQty(request.Context(), cartUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *ShoppingCartControllerImpl) DeleteCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := web.Claims{}
	helper.ParseJwt(request, &claim)

	cartsDeleteRequest := web.CartsDeleteRequest{}
	helper.ReadFromRequestBody(request, &cartsDeleteRequest)

	carts := []web.CartDeleteRequest{}

	for _, cart := range cartsDeleteRequest.Detail {
		cart.UserId = claim.Id
		carts = append(carts, cart)
	}

	cartDelete := web.CartsDeleteRequest{
		Detail: carts,
	}
	controller.Service.DeleteCart(request.Context(), cartDelete.Detail)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) FindSome(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartRequest := web.CartsSelectRequest{}
	helper.ReadFromRequestBody(request, &cartRequest)

	cartResponse := controller.Service.FindSomeCart(request.Context(), cartRequest.Detail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
