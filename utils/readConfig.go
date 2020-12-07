package utils

import (
	"github.com/spf13/viper"
	"log"
)

func ReadConfig(key string) string {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	value := viper.GetString(key)
	return value
}
