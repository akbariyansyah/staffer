package config

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
		ServerHost:       "localhost",
		ServerPort:       "4000",
		DatabaseHost:     "localhost",
		DatabasePort:     "5432",
		DatabaseUser:     "postgres",
		DatabasePassword: "P@ssW02d123",
		DatabaseName:     "db_worker",
	}
}
func NewConfig() *Config {
	return &conf
}
