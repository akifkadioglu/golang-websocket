package utils

import (
	"fmt"
	"net/http"

	"github.com/akifkadioglu/golang-websocket/env"
	"github.com/akifkadioglu/golang-websocket/models"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

/* tokenAuth'u .env dosyasındaki appkey ile anlamlandırıyor */
func InitToken() {
	tokenAuth = jwtauth.New("HS256", []byte(env.Getenv(env.APP_KEY)), nil)
}

/* getter setter gibi düşünebilirsin */
func JWTAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

/* jwt'yi döndürüyor */
func GetUser(r *http.Request) models.JwtModel {
	_, claims, _ := jwtauth.FromContext(r.Context())
	fmt.Println(claims)
	var user models.JwtModel
	MapToStruct(claims, &user)

	return user
}
