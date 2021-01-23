package impl

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/gonduit/users"
)

type Router struct {
	E       *echo.Echo
	RegServ *users.RegisterService
}

func NewRouter(e *echo.Echo, regServ *users.RegisterService) *Router {
	r := &Router{
		E:       e,
		RegServ: regServ,
	}
	return r
}

func (r *Router) SetupEndpoints() error {
	r.E.GET("/greetings/:user", Greeting)
	r.E.POST("/users", RegisterHandler{Service: r.RegServ}.Register)

	return nil
}

func Greeting(c echo.Context) error {
	tester := c.Param("user")
	return c.String(http.StatusOK, "hello, "+tester+"!")
}
