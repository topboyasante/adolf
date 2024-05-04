package main

import (
	"github.com/topboyasante/adolf/internal/actions"
	"github.com/topboyasante/adolf/internal/templates"
	"github.com/urfave/cli"
	"log"
	"os"
)

func initAction(c *cli.Context) error {
	folders := []string{
		"cmd/main",
		"internal/config",
		"internal/routes",
		"internal/models",
		"internal/controllers",
		"internal/utils",
	}
	files := []string{
		"cmd/main/main.go",
		"internal/config/app.go",
		"internal/utils/utils.go",
		"internal/routes/routes.go",
		"internal/models/demo_model.go",
		"internal/controllers/demo_controller.go",
	}
	modelTemplate := templates.GenerateModelTemplate()
	routeTemplate := templates.GenerateRouteTemplate()
	configTemplate := templates.GenerateDBConfigTemplate()
	mainTemplate := templates.GenerateMainTemplate()
	utilsTemplate := templates.GenerateUtilsemplate()

	actions.InitializeSetup()

	for _, value := range folders {
		err := os.MkdirAll(value, 0750)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, value := range files {
		switch value {
		case "cmd/main/main.go":
			{
				actions.CreateFile(value, mainTemplate, 0750)
			}
		case "internal/models/demo_model.go":
			{
				actions.CreateFile(value, modelTemplate, 0750)
			}
		case "internal/config/app.go":
			{
				actions.CreateFile(value, configTemplate, 0750)
			}
		case "internal/utils/utils.go":
			{
				actions.CreateFile(value, utilsTemplate, 0750)
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
