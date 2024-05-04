package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
	"github.com/topboyasante/adolf/internal/actions"
	"github.com/topboyasante/adolf/internal/templates"
	"github.com/urfave/cli"
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
	configTemplate := templates.GenerateDBConfigDefaultTemplate()
	mainTemplate := templates.GenerateMainTemplate()
	utilsTemplate := templates.GenerateUtilsTemplate()
	controllerTemplate := templates.GenerateControllerTemplate()

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
		case "internal/controllers/demo_controller.go":
			{
				actions.CreateFile(value, controllerTemplate, 0750)
			}
		default:
			{
				actions.CreateFile(value, "package main", 0750)
			}
		}
	}

	return nil
}

func readAdolfConfig(filepath string) string {
	dat, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal("Could not find config.adolf.toml file")
		os.Exit(1)
	}

	return string(dat)
}

func GenerateWithConfigAction(c *cli.Context) {

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
	mainTemplate := templates.GenerateMainTemplate()
	utilsTemplate := templates.GenerateUtilsTemplate()
	controllerTemplate := templates.GenerateControllerTemplate()

	configFilePath := c.Args().Get(2)

	configFile := readAdolfConfig(configFilePath)

	var cfg templates.AdolfDBConfig
	err := toml.Unmarshal([]byte(configFile), &cfg)
	if err != nil {
		log.Fatal("Failed to load toml file: ", err)
	}

	// TODO poll the variables and generate database config from that.
	configTemplate := templates.GenerateDBConfigTemplate(cfg)

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
		case "internal/controllers/demo_controller.go":
			{
				actions.CreateFile(value, controllerTemplate, 0750)
			}
		default:
			{
				actions.CreateFile(value, "package main", 0750)
			}
		}
	}

}
