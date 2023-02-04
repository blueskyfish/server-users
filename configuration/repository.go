package configuration

import (
	"bytes"
	"github.com/blueskyfish/server-users/repository"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
)

const dsnTemplate = "{{ .User }}:{{ .Password }}@tcp({{ .Host }}:{{ .Port }})/{{ .Database }}?charset=utf8mb4&parseTime=True&loc=Local"

type dbConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (conf Configuration) Repository() *repository.Repository {
	if conf.repository != nil {
		return conf.repository
	}
	logger := conf.WithGroup("repository")
	dbConfig := dbConfig{}
	err := viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		logger.Errorf("missing database configuration %v", err)
	}
	dsn, err := parseDbConfig(dbConfig)
	if err != nil {
		logger.Errorf("missing database configuration %v", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("could not open database %v", err)
	}

	repo := repository.NewRepository(db, logger)

	conf.repository = repo

	return repo
}

func parseDbConfig(dbConfig dbConfig) (string, error) {
	tmpl, err := template.New("db-config").Parse(dsnTemplate)
	if err != nil {
		return "", err
	}
	buffer := bytes.Buffer{}
	err = tmpl.Execute(&buffer, dbConfig)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
