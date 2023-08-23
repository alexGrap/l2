package main

import (
	"L2/devTasks/Task11/config"
	"L2/devTasks/Task11/internal/delivery"
)

func main() {
	cfg := config.GetConfigYml()
	app := delivery.Fabric()
	err := app.Hearing(cfg)
	if err != nil {
		return
	}
}
