package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "short",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "server", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
					&cli.BoolFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().Get(0))
					fmt.Println("server:", c.Bool("server"))
					fmt.Println("option:", c.Bool("option"))
					fmt.Println("message:", c.Bool("message"))
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
