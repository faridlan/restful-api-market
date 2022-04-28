package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type ProductDetailControllerImpl struct {
	Service service.ProductDetailService
}

func NewProductDetailController(service service.ProductDetailService) ProductDetailController {
	return &ProductDetailControllerImpl{
		Service: service,
	}
}

func (controller *ProductDetailControllerImpl) FindProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.Service.FindProduct(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *ProductDetailControllerImpl) AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

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
