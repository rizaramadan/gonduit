package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/gonduit/users"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test/:tester", func(c echo.Context) error {
		tester := c.Param("tester")
		return c.String(http.StatusOK, "hello, "+tester+"!")
	})

	e.POST("/users", func(c echo.Context) (err error) {
		i := new(users.RegisterInputwrapper)
		if err = c.Bind(i); err != nil {
			return err
		}
		u := i.GetUser()
		o := users.GetUserOutput(u)
		return c.JSON(http.StatusOK, o)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
