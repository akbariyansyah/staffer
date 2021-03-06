package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDatabase(conf *Config) (*sql.DB, error) {

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.DB_USER, conf.DB_PASS, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
