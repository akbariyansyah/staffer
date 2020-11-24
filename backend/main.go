package main

import (
	"backend/config"
	"backend/database"
	"backend/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.New()
	c := config.NewConfig()
	//fmt.Println(os.Getenv("SERVER_HOST"))
	//fmt.Println(os.Getenv("SERVER_PORT"))
	//fmt.Println(os.Getenv("DATABASE_HOST"))
	//fmt.Println(os.Getenv("DATABASE_PORT"))
	//fmt.Println(os.Getenv("DATABASE_USER"))
	//fmt.Println(os.Getenv("DATABASE_PASSWORD"))
	//fmt.Println(os.Getenv("DATABASE_NAME"))
	db := database.StartDatabase(c)

	router.NewRouter(g,db)

	log.Println("Server its starting at port 4000")


	g.Run(c.ServerHost + ":" + c.ServerPort)
}