package app

import (
	"github.com/urfave/cli"
)

// func generate will return a command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "command line application"
	app.Usage = "search IPS and server names on the internet"
	return app
}
