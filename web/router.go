package web

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizaramadan/gonduit/contracts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	E          *echo.Echo
	regHandler *RegisterHandler
}

func NewRouter(e *echo.Echo, regHandler *RegisterHandler) *Router {
	r := &Router{
		E:          e,
		regHandler: regHandler,
	}
	return r
}

func (r *Router) SetupEndpoints() error {
	jwtAuth := newJwtMiddleware()
	r.E.GET("/greetings/:user", greeting)
	r.E.POST("/users", r.regHandler.Register)
	r.E.GET("/welcome", jwtAuth(welcome))

	return nil
}

func newJwtMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(contracts.Secret),
	}
	jwtAuth := middleware.JWTWithConfig(config)
	return jwtAuth
}

func greeting(c echo.Context) error {
	tester := c.Param("user")
	return c.String(http.StatusOK, "hello, "+tester+"!")
}

func welcome(c echo.Context) error {
	username, _, err := getUsernameEmail(c)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Welcome "+username+"!")
}
