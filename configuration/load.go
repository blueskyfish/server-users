package configuration

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("users-config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		// Logger
		log.Debug("Load configuration error", err, "\n")
	}
}
