package database

import (
	"backend/config"
	"context"
	"log"

	"github.com/go-pg/pg"
)

func StartDatabase(conf *config.Config) *pg.DB {
	
	db := pg.Connect(&pg.Options{
		Addr:     conf.DatabaseHost + ":" + conf.DatabasePort,
		User:     conf.DatabaseUser,
		Password: conf.DatabasePassword,
		Database: conf.DatabaseName,
	})
	if _, err := db.ExecContext(context.Background(),"SELECT 1");err != nil {
		log.Fatal(err)
	}
	log.Println("Starting database postgres ")

	return db
}
