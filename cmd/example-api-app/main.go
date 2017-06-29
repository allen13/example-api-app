package main

import (
	"fmt"
	"log"

	"github.com/allen13/example-api-app/pkg/api"
	"github.com/allen13/example-api-app/pkg/command"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Golang App Standard.

		Usage:
		  example-api-app (api|command) [--config=<config>]
		  example-api-app -h | --help
		  example-api-app --version

		Options:
		  -h --help     Show this screen.
		  --version     Show version.
		  --config=<config>  Config file [default: /etc/golang-app-standard/golang-app-standard.toml].`

	args, err := docopt.Parse(usage, nil, true, "Golang App Standard 1.0", false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(args)

	if args["command"].(bool) {
		command.RunCommand()
	}

	if args["api"].(bool) {
		api.StartAPI()
	}
}
