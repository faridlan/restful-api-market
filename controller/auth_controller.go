package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	MyProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateImg(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
