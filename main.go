package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizaramadan/gonduit/impl/persistence"
	"github.com/rizaramadan/gonduit/users"
	"github.com/rizaramadan/gonduit/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	dsn := "host=localhost user=postgres password=mainmain dbname=gonduit port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot access db")
	}
	regHandler := setupRegisterHandler(db)

	r := web.NewRouter(e, regHandler)
	if err := r.SetupEndpoints(); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func setupRegisterHandler(db *gorm.DB) *web.RegisterHandler {
	userRepo := persistence.NewUserRepoDb(db)
	regServ := users.NewRegisterService(userRepo)
	regHandler := new(web.RegisterHandler)
	regHandler.Service = regServ
	return regHandler
}
