package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "hello",
		Action: func(_ *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "greet",
				Action: func(_ *cli.Context) error {
					if len(os.Args) < 3 {
						return errors.New("name is required")
					}

					fmt.Println("Hello, " + os.Args[2] + "!")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
