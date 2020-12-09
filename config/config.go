package config

// Config -> basic database config type
type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func NewConfig() *Config {
	conf := &Config{}
	conf.DB_HOST = "localhost"
	conf.DB_PORT = "3306"
	conf.DB_USER = "root"
	conf.DB_PASS = "password123"
	conf.DB_NAME = "worker"

	//conf.DB_HOST = os.Getenv("DB_HOST")
	//conf.DB_PORT = os.Getenv("DB_PORT")
	//conf.DB_USER = os.Getenv("DB_USER")
	//conf.DB_PASS = os.Getenv("DB_PASS")
	//conf.DB_NAME = os.Getenv("DB_NAME")

	//switch {
	//case os.Getenv("DB_HOST") == "":
	//	conf.DB_HOST = utils.ReadConfig("database.host")
	//case os.Getenv("DB_PORT") == "":
	//	conf.DB_PORT = utils.ReadConfig("database.port")
	//case os.Getenv("DB_USER") == "":
	//	conf.DB_USER = utils.ReadConfig("database.user")
	//case os.Getenv("DB_PASS") == "":
	//	conf.DB_PASS = utils.ReadConfig("database.pass")
	//case os.Getenv("DB_NAME") == "":
	//	conf.DB_NAME = utils.ReadConfig("database.name")
	//}

	return conf
}
