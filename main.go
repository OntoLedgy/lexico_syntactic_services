package main

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/code/infrastructure/cli"
)

var (
	// These fields are populated by govvv
	Version    string
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
)

func main() {

	fmt.Printf("Starting Syntactic Checker Version %s-%s-%s\n", Version, GitCommit, BuildDate)
	cli.Initialise()

}
