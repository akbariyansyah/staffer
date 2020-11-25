package config

import "os"

type Config struct {
	ServerHost       string
	ServerPort       string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

var conf Config

func init() {
	conf = Config{
		ServerHost:       os.Getenv("SERVER_HOST"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
	}
}
func NewConfig() *Config {
	return &conf
}
