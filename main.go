package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"staffer/config"
	"staffer/utils"
)

func main() {
	e := echo.New()

	conf := config.NewConfig()
	db := config.NewDatabase(conf)
	fmt.Println(db)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello there")
	})

	address := os.Getenv("SERVER_PORT")
	if address == "" {
		address = utils.ReadConfig("server.port")
	}
}
