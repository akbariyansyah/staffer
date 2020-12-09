package main

import (
	"os"
	"staffer/api"
	"staffer/config"
	"staffer/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	conf := config.NewConfig()
	db, err := config.NewDatabase(conf)
	if err != nil {
		panic(err)
	}
	api.NewRoutes(e, db)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	address := os.Getenv("SERVER_PORT")
	if address == "" {
		address = utils.ReadConfig("server.port")
	}
	e.Logger.Fatal(e.Start("localhost:" + address))
	//e.Logger.Fatal(e.Start(":" + address))
}
