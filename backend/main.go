package main

import (
	"backend/config"
	"backend/database"
	"backend/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.Default()
	c := config.NewConfig()
	db := database.StartDatabase(c)
	router.NewRouter(g,db)

	log.Println("Server its starting at port 4000")

	g.Run(c.ServerHost + ":" + c.ServerPort)
}