package helper

import (
	"net/http"
	"strings"

	"github.com/faridlan/restful-api-market/model/web"
	"github.com/golang-jwt/jwt/v4"
)

func ParseJwt(request *http.Request, claim *web.Claims) web.Claims {
	authorizationHeader := request.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	_, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return web.JwtSecret, nil
	})
	PanicIfError(err)

	return *claim
}
