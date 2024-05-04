package main

import (
	"fmt"
	"log"
	"os"

	"github.com/topboyasante/adolf/internal/actions"
	"github.com/topboyasante/adolf/internal/templates"
	"github.com/urfave/cli"
)

func initAction(c *cli.Context) error {
	//List of folders that should be created
	folders := []string{"cmd/main", "internal/routes", "internal/models"}
	//list of files that should be created
	files := []string{"cmd/main/main.go", "internal/routes/routes.go", "internal/models/demo_model.go"}

	modelTemplate := templates.GenerateModelTemplate()
	routeTemplate := templates.GenerateRouteTemplate()

	fmt.Println("Bootstrapping your web application...")

	for _, value := range folders {
		createFolderError := os.MkdirAll(value, 0750)
		if createFolderError != nil {
			log.Fatal(createFolderError)
		}
	}

	for _, value := range files {
		switch value {
		case "internal/models/demo_model.go":
			{
				actions.CreateFile(value, modelTemplate, 0750)
			}
		case "internal/routes/routes.go":
			{
				actions.CreateFile(value, routeTemplate, 0750)
			}
		default:
			{
				actions.CreateFile(value, "package main", 0750)
			}
		}
	}

	return nil
}
