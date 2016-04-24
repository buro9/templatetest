package views

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	doFlags      sync.Once
	templatePath string
)

func RegisterFlags() {
	doFlags.Do(func() {
		flag.StringVar(&templatePath, "templatePath", "templates", "Path to the templates")

		if !strings.HasSuffix(templatePath, "/") {
			templatePath += "/"
		}

		if _, err := os.Stat(templatePath); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("path to templates '%s' does not exist: %s", templatePath, err.Error())
			}
			if os.IsPermission(err) {
				log.Fatalf("no permission to path to templates '%s': %s", templatePath, err.Error())
			}
			log.Fatal(err)
		}
	})
}
