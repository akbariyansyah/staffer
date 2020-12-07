package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"staffer/config"

	"staffer/utils"
)

func main() {
	e := echo.New()

	conf := config.NewConfig()
	db := config.NewDatabase(conf)

	config.NewRoutes(e, db)

	address := os.Getenv("SERVER_PORT")
	if address == "" {
		address = utils.ReadConfig("server.port")
	}
	e.Logger.Fatal(e.Start(":"+address))
}
