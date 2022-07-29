package controller

import (
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productUpdateRequest := web.ProductUpateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := params.ByName("productId")

	productUpdateRequest.IdProduct = productId

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	controller.ProductService.Delete(request.Context(), productId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	productResponse := controller.ProductService.FindyId(request.Context(), productId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pagination := helper.Pagination(request)
	productResponses := controller.ProductService.FindAll(request.Context(), pagination)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) CreateImg(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	file, fileHeader, err := request.FormFile("productImage")
	helper.PanicIfError(err)
	defer file.Close()

	// fileBytes, err := ioutil.ReadAll(file)
	// helper.PanicIfError(err)

	// image := web.ProductCreateRequest{
	// 	ImageUrl: string(fileBytes),
	// }

	storage := domain.Storage{
		Name: fileHeader.Filename,
		File: file,
	}

	productResponse := controller.ProductService.CreateImg(request.Context(), storage)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
