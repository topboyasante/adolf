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
	app.UsageText=""

	// Define the 'init' command
	app.Commands = []cli.Command{
		{
			Name:      "init",
			Usage:     "Initializes a go backend application with MySQL as the DB Provider",
			Action:    initAction,
			UsageText: `Run "adolf init". A go backend application will be set up with a mySQL driver`,
		},
		{
			Name:      "heil",
			Usage:     "Initializes a go backend application with the DB Provider you specify in your .adolf.toml file",
			Action:    GenerateWithConfigAction,
			UsageText: `Run "adolf heil config.adolf.toml". A go backend application will be set up with a driver based on the config in your .toml file`,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
