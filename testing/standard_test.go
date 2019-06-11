package testing

import (
	"fmt"
	"os"
	"syntactic_checker/orchestrators"
	"testing"
)

func TestStandard(t *testing.T) {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	configuration_file_path := os.Args[3]

	fmt.Printf("configuration file path:%s\n", os.Args[3])

	orchestrators.
		Orchestrate_syntactic_checks(
			configuration_file_path)

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}
