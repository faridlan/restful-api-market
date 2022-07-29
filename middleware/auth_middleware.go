package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model"
	"github.com/faridlan/restful-api-market/model/web"
	"github.com/faridlan/restful-api-market/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler    http.Handler
	Repository repository.BlacklistRepository
	DB         *sql.DB
}

func NewAuthMiddleware(handler http.Handler, repository repository.BlacklistRepository, db *sql.DB) *AuthMiddleware {
	return &AuthMiddleware{
		Handler:    handler,
		Repository: repository,
		DB:         db,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authorizationHeader := request.Header.Get("Authorization")

	if request.URL.Path == "/api/login" || request.URL.Path == "/api/register" || request.URL.Path == "/" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	tx, err := middleware.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbak(tx)

	_, err = middleware.Repository.SelectById(request.Context(), tx, tokenString)

	if err == nil {
		//UNAUTHORIZED 401
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	} else {
		//OK 200
		var claim = &web.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != web.JwtSigningMEethod {
				return nil, fmt.Errorf("signing method invalid")
			}
			return web.JwtSecret, nil
		})

		if err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   err.Error(),
			}

			helper.WriteToResponseBody(writer, webResponse)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
			}

			helper.WriteToResponseBody(writer, webResponse)
			writer.WriteHeader(http.StatusBadRequest)
			return
		} else {

			endpoints := model.Endpoints()
			for _, e := range endpoints {
				if request.URL.Path == e.Url && request.Method == e.Method && claim.RoleId != 1 {
					writer.Header().Add("Content-Type", "application/json")
					writer.WriteHeader(http.StatusUnauthorized)
					webResponse := web.WebResponse{
						Code:   http.StatusUnauthorized,
						Status: "UNAUTHORIZED",
					}
					helper.WriteToResponseBody(writer, webResponse)
					return
				}
			}

			middleware.Handler.ServeHTTP(writer, request)

		}
	}

}
