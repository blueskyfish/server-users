package configuration

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	name string
	log  *logrus.Logger
}

const StageDev = "dev"
const StateProd = "prod"

func NewConfiguration() Configuration {
	conf := Configuration{
		name: "",
		log:  nil,
	}
	level := viper.GetString("log.level")
	conf.log = logrus.New()
	conf.log.SetFormatter(&logrus.JSONFormatter{})
	conf.log.SetOutput(os.Stdout)

	if level != "" {
		logLevel, err := logrus.ParseLevel(level)
		if err != nil {
			conf.log.SetLevel(logrus.DebugLevel)
		} else {
			conf.log.SetLevel(logLevel)
		}
	}
	return conf
}
