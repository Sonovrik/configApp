package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Title  string `yaml:"title"`
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	DB     DB     `yaml:"postgres"`
}

type Server struct {
	Bind string `yaml:"bind"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
}

func Init(pathToConfig string) (app *AppConfig, e error) {
	yamlFile, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, errors.New("Cannot open config file " + pathToConfig + " : " + err.Error())
	}

	if err = yaml.Unmarshal(yamlFile, app); err != nil {
		return nil, errors.New("Cannot parse yaml config file " + pathToConfig + " : " + err.Error())
	}

	return app, nil
}
