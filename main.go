package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizaramadan/gonduit/impl"
	"github.com/rizaramadan/gonduit/users"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	r := impl.NewRouter(e, &users.RegisterService{})
	if err := r.SetupEndpoints(); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
