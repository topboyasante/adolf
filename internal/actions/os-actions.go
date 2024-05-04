package actions

import (
	"io/fs"
	"log"
	"os"
)

func CreateFile(path string, template string, perm fs.FileMode) {
	err := os.WriteFile(path, []byte(template), perm)
	if err != nil {
		log.Fatal(err)
	}
}
