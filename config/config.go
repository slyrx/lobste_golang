// config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	Config = v
}
