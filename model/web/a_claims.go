package web

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSigningMEethod = jwt.SigningMethodHS256
var JwtSecret = []byte("Anj1ngAd4l4hH3w4nuuAAii8sdA73DFed7")
var ExpiredTime = time.Now().Add(time.Hour * 24)

type Claims struct {
	Id       int    `json:"id,omitempty"`
	IdUser   string `json:"id_user,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	RoleId   int    `json:"role_id,omitempty"`
	Token    string `json:"token,omitempty"`
	*jwt.RegisteredClaims
}
