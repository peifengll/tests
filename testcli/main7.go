package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Value:   "english",
				Usage:   "balabalabala",
				EnvVars: []string{"app_lang"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("lang") == "english" {
				fmt.Println("hello ")
			} else {
				fmt.Println("你好")
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/**
app_lang=k ./env

*/
