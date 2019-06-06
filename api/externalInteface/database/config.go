package database

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB   DBConf
	Addr string `envconfig:"port" default:":8080"`
}

type DBConf struct {
	User     string `envconfig:"mysql_user" default:"root"`
	Password string `envconfig:"mysql_password" default:"accelia"`
	Host     string `envconfig:"mysql_ip" default:"10.0.72.20:32484"`
	DB       string `envconfig:"mysql_db" default:"gomi"`
}

func LoadConfig() (*Config, error) {
	var config = Config{}
	if err := envconfig.Process("APP", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
