package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizaramadan/gonduit/impl"
	"github.com/rizaramadan/gonduit/persistence"
	"github.com/rizaramadan/gonduit/users"
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

	r := impl.NewRouter(e, regHandler)
	if err := r.SetupEndpoints(); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func setupRegisterHandler(db *gorm.DB) *impl.RegisterHandler {
	userRepo := persistence.NewUserRepoDb(db)
	regServ := users.NewRegisterService(userRepo)
	regHandler := new(impl.RegisterHandler)
	regHandler.Service = regServ
	return regHandler
}
