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
		{
			Name:   "heil",
			Usage:  "Initialize the web application with a config.adolf.toml",
			Action: initAction, // TODO separate to different method
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
