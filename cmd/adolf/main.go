package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Adolf"
	app.Usage = "Initialize a go web application"

	// Define the 'init' command
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize the web application",
			Action: initAction,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
