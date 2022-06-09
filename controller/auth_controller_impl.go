package controller

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authServie service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authServie,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse, claims := controller.AuthService.Register(request.Context(), userCreateRequest)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(web.JwtSecret)
	helper.PanicIfError(err)

	userResponse.Token = tokenString

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) CreateUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.AuthService.CreateUsers(request.Context(), userCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginCreateRequest := web.LoginCreateRequest{}
	helper.ReadFromRequestBody(request, &loginCreateRequest)

	user, claims := controller.AuthService.Login(request.Context(), loginCreateRequest)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(web.JwtSecret)
	helper.PanicIfError(err)

	user.Token = tokenString

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) MyProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := &web.Claims{}
	claims := helper.ParseJwt(request, claim)

	userResponse := controller.AuthService.Profile(request.Context(), claims.IdUser)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller AuthControllerImpl) Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	userResponse := controller.AuthService.Profile(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) UpdateProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")

	userUpdateRequest.IdUser = userId

	userResponse := controller.AuthService.UpdateProfile(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pagination := helper.Pagination(request)
	userResponses := controller.AuthService.FindAll(request.Context(), pagination)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	authHeader := request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	blacklistCreate := web.BlacklistCreateRequest{
		Token: tokenString,
	}

	controller.AuthService.Logout(request.Context(), blacklistCreate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) CreateImg(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	file, _, err := request.FormFile("profileImage")
	helper.PanicIfError(err)
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	helper.PanicIfError(err)

	image := web.UserCreateRequest{
		ImageUrl: string(fileBytes),
	}

	userResponse := controller.AuthService.UploadImage(request.Context(), image)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
