package web

import (
	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/gonduit/users"
	"net/http"
)

type RegisterHandler struct {
	Service *users.RegisterService
}

//Register will handle POST /users/ endpoint
func (r *RegisterHandler) Register(c echo.Context) (err error) {
	i := new(users.RegisterInputWrapper)

	if err = c.Bind(i); err != nil {
		return err
	}

	if err = i.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid register input")
	}

	o, err := r.Service.Register(i)
	if err != nil {
		return err
	}

	o.User.Token, err = GenerateToken(i.GetUser())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, o)
}
