package main

import (
	"configApp/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	var configPath string = "configs/config.yaml"

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	logrus.Info(cfg.Title)
}
