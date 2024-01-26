package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

// 测试参数的获取和输出

func main() {
	app := &cli.App{
		Name:  "测试参数",
		Usage: "测试一次参数获取",
		Action: func(c *cli.Context) error {
			fmt.Println(c.NArg(), ":")
			list := c.Args()
			for k, v := range list.Slice() {
				fmt.Println(k, v)
			}
			fmt.Println(c.Args())
			return nil
		},
	}
	app.Run(os.Args)
}
