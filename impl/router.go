package impl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	E          *echo.Echo
	regHandler *RegisterHandler
}

func NewRouter(e *echo.Echo, regHndlr *RegisterHandler) *Router {
	r := &Router{
		E:          e,
		regHandler: regHndlr,
	}
	return r
}

func (r *Router) SetupEndpoints() error {
	r.E.GET("/greetings/:user", Greeting)
	r.E.POST("/users", r.regHandler.Register)

	return nil
}

func Greeting(c echo.Context) error {
	tester := c.Param("user")
	return c.String(http.StatusOK, "hello, "+tester+"!")
}
