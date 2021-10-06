package handlers

import "github.com/dgrijalva/jwt-go"

type AuthHandler struct{}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
