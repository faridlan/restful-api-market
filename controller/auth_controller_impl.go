package controller

import (
	"net/http"
	"strconv"
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

	userResponse := controller.AuthService.Register(request.Context(), userCreateRequest)

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

	claims := controller.AuthService.Login(request.Context(), loginCreateRequest)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(web.JwtSecret)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tokenString,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claim := &web.Claims{}
	claims := helper.ParseJwt(request, claim)

	userResponse := controller.AuthService.Profile(request.Context(), claims.Id)
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
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.AuthService.UpdateProfile(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userResponses := controller.AuthService.FindAll(request.Context())
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
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
