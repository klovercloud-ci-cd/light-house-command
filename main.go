package main

import (
	"github.com/klovercloud-ci-cd/light-house-command/api"
	"github.com/klovercloud-ci-cd/light-house-command/config"
	_ "github.com/klovercloud-ci-cd/light-house-command/docs"
)

// @title Klovercloud-ci-light-house-command API
// @description Klovercloud-light-house-command API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
