package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/OntoLedgy/syntactic_checker/code/services/service_run_preparers"
	"github.com/OntoLedgy/syntactic_checker/code/services/syntactic_checking_services"
)

func TestStandard(t *testing.T) {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	fmt.Print(
		"Starting syntactic checking service\n")

	configuration_file_path :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\test_data_20190918\\citadel_document_types\\syntactic_checker_configuration.json"

	service_run_preparer :=
		new(
			service_run_preparers.ServiceRunPreparers)

	service_run_data :=
		service_run_preparer.
			Get_service_run_data(
				configuration_file_path)

	syntactic_checking_service_factory :=
		new(
			syntactic_checking_services.
				SyntacticCheckingServiceFactory)

	syntactic_checking_service :=
		syntactic_checking_service_factory.
			Create(
				service_run_data)

	syntactic_checking_service.
		Run_syntactic_checking_service()

	fmt.Print(
		"Ending syntactic checking service")

	//logging.End_logger()

	trace()

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}

func trace() {

	pc := make([]uintptr, 10) // at least 1 entry needed

	runtime.Callers(2, pc)

	f := runtime.FuncForPC(pc[0])

	file, line := f.FileLine(pc[0])

	function_name := f.Name()

	fmt.Printf("file: %s:\nfunction name: %s\nline: %d\nfunction entry: %d\n", file, function_name, line, f.Entry())
}
