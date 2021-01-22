package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test/:tester", func(c echo.Context) error {
		tester := c.Param("tester")
		return c.String(http.StatusOK, "hello, " + tester +"!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
