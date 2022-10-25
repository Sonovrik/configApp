package config

import (
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
		return nil, err
	}

	app = &AppConfig{}
	if err = yaml.Unmarshal(yamlFile, app); err != nil {
		return nil, err
	}

	return app, nil
}
