package configuration

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func (conf Configuration) Name() string {
	if conf.name == "" {
		conf.name = viper.GetString("name")
	}
	return conf.name
}

func (conf Configuration) Address() string {
	port := viper.GetInt("http.port")
	if port == 0 {
		port = 8080
	}
	host := viper.GetString("http.host")
	if host == "" {
		host = "localhost"
	}
	return host + ":" + strconv.Itoa(port)
}

func (conf Configuration) Log() *logrus.Logger {
	return conf.log
}

func (conf Configuration) Stage() string {
	state := viper.GetString("stage")
	if state == "" {
		state = "dev"
	}
	return state
}
