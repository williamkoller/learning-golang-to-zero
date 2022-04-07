package main

import (
	"log"
	"os"

	"github.com/williamkoller/command-line/app"
)

func main() {
	application := app.Generate()

	error := application.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}