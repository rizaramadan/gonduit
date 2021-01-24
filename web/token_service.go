package web

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/gonduit/contracts"
	"github.com/rizaramadan/gonduit/users"
	"time"
)

func GenerateToken(u users.User) (string, error) {
	if len(u.Username) <= 0 && len(u.Email) <= 0 {
		return "", errors.New("invalid username and password for generate token")
	}
	return generateToken(u)
}

func generateToken(u users.User) (string, error) {
	claims := &jwtCustomClaims{
		u.Username,
		u.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(contracts.Secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func getUsernameEmail(c echo.Context) (string, string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	username := claims.Username
	email := claims.Email
	return username, email, nil
}
