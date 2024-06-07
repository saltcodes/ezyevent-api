package util

import (
	"github.com/spf13/viper"
	"log"
)

func GetVariableWith(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invalid Key")
	}
	return value
}
