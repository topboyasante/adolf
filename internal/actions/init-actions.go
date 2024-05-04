package actions

import (
	"fmt"
	"os"
	"os/exec"
)

func InitializeSetup() {
	var moduleName string

	fmt.Println("What is the name of your module?")
	fmt.Scanln(&moduleName)

	cmd := exec.Command("go", "mod", "init", moduleName)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error running 'go mod init':", err)
		os.Exit(1)
	}

	fmt.Println("Bootstrapping your web application...")

}
