package main

import (
	"github.com/klovercloud-ci-cd/light-house-command/api"
	"github.com/klovercloud-ci-cd/light-house-command/config"
)

func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
