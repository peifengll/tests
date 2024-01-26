package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Value: 0,
				Usage: "use a randomized port",
				//DefaultText: "random",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("err")
	}
}
