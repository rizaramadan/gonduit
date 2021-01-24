package web

import "github.com/dgrijalva/jwt-go"

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
