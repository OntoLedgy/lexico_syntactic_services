package cli

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/code/services"
	"github.com/OntoLedgy/syntactic_checker/code/version"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func Initialise() {

	info()
	commands()

	app.Action = noArgs

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

func info() {
	app.Name = "Syntactic Checker"
	app.Usage = "Syntactic Checking application. "
	app.Author = "khanm"
	app.Version = version.BuildVersion
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "identified_string_list_checks",
			Aliases: []string{"i"},
			Usage:   "Checks a list of identified strings using the checks types in the configuration file",
			Action: func(c *cli.Context) {
				services.Orchestrate_services(c)
			},
		},
		{
			Name:    "column_checks",
			Aliases: []string{"c"},
			Usage:   "To be implemented",
			Action: func(c *cli.Context) {
				fmt.Printf("to be implenmented...\n")
				fmt.Printf("version : %s\n", version.BuildVersion)
			},
		},
		{
			Name:    "table_checks",
			Aliases: []string{"t"},
			Usage:   "To be implemented",
			Action: func(c *cli.Context) {
				fmt.Printf("to be implenmented...")
			},
		},
	}
}

func noArgs(c *cli.Context) {

	services.Orchestrate_services(c)
}
