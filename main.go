package main

import (
	"github.com/klovercloud/lighthouse-command/api"
	"github.com/klovercloud/lighthouse-command/config"
)

func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
