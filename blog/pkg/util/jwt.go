package util

import (
	"blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
    jwt.StandardClaims
}