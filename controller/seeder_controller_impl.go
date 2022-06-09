package controller

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/domain"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/julienschmidt/httprouter"
)

type SeederControllerImpl struct {
	SeederService  service.SeederService
	UserService    service.AuthService
	AddressService service.AddressService
	ProductService service.ProductService
	OrderService   service.ShippingAddressService
}

func NewSeederController(SeederService service.SeederService, UserService service.AuthService, AddressService service.AddressService, ProductService service.ProductService, OrderService service.ShippingAddressService) SeederController {
	return &SeederControllerImpl{
		SeederService:  SeederService,
		UserService:    UserService,
		ProductService: ProductService,
		AddressService: AddressService,
		OrderService:   OrderService,
	}
}

//go:embed json/user.json
//go:embed json/product.json
//go:embed json/address.json

var Json embed.FS

func (controller *SeederControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	//product
	product, err := Json.ReadFile("json/product.json")
	helper.PanicIfError(err)
	productCreate := []web.ProductCreateRequest{}
	err = json.Unmarshal(product, &productCreate)
	helper.PanicIfError(err)

	for _, v := range productCreate {
		controller.ProductService.Create(request.Context(), v)
	}

	//user
	user, err := Json.ReadFile("json/user.json")
	helper.PanicIfError(err)
	userCreate := []web.UserCreateRequest{}
	err = json.Unmarshal(user, &userCreate)
	helper.PanicIfError(err)

	for _, v := range userCreate {
		controller.UserService.Register(request.Context(), v)
	}

	//address
	address, err := Json.ReadFile("json/address.json")
	helper.PanicIfError(err)
	addressCreate := []web.AddressCreateRequest{}
	err = json.Unmarshal(address, &addressCreate)
	helper.PanicIfError(err)

	for i, v := range addressCreate {
		pagintaion := domain.Pagination{
			Page:  i + 1,
			Limit: i + 1,
		}
		userResponse := controller.UserService.FindSeeder(request.Context(), pagintaion)
		v.UserId = userResponse.Id
		controller.AddressService.Create(request.Context(), v)
	}

	//orders
	for i, v := range userCreate {
		log.Print(v)
		pagintaion := domain.Pagination{
			Page:  i,
			Limit: i + 1,
		}
		pagintaionUser := domain.Pagination{
			Page:  i + 1,
			Limit: i + 1,
		}

		pagintaionProducts := domain.Pagination{
			Page:  i + 3,
			Limit: 2,
		}
		userResponse := controller.UserService.FindSeeder(request.Context(), pagintaionUser)
		address := controller.AddressService.FindSeeder(request.Context(), pagintaion)
		productsResponse := controller.ProductService.FindSeeder(request.Context(), pagintaionProducts)

		y := []web.CreateOrder{}
		for _, v := range productsResponse {
			x := web.CreateOrder{
				UserId:    userResponse.Id,
				ProductId: v.IdProduct,
				Quantity:  3,
			}

			y = append(y, x)
		}

		x := web.OrderCreateRequest{
			UserId:    userResponse.Id,
			AddressId: address.IdAddress,
			Products:  y,
		}

		controller.OrderService.CreateOrder(request.Context(), x)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SeederControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.SeederService.Delete(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
