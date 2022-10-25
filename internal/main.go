package main

import (
	"log"

	"configApp/internal/config"
)

func main() {
	var configPath string = "config.yaml"
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err.Error())
	}
}
