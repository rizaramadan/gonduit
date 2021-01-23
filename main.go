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

	regHandler := setupRegisterHandler()

	r := impl.NewRouter(e, regHandler)
	if err := r.SetupEndpoints(); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func setupRegisterHandler() *impl.RegisterHandler {
	regServ := new(users.RegisterService)
	regHandler := new(impl.RegisterHandler)
	regHandler.Service = regServ
	return regHandler
}
