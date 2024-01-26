package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var language string

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "english" {
				fmt.Println("hello", name)
			} else {
				fmt.Println("你好", name)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
