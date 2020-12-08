package main

import (
	"os"
	"staffer/api"
	"staffer/config"
	"staffer/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	conf := config.NewConfig()
	db,err := config.NewDatabase(conf)
	if err != nil {
		panic(err)
	}
	api.NewRoutes(e,db)

	address := os.Getenv("SERVER_PORT")
	if address == "" {
		address = utils.ReadConfig("server.port")
	}
	e.Logger.Fatal(e.Start("localhost:" + address))
}
