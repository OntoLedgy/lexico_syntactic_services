package syntactic_check_service

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/services/cell_list_checks_service"
	"syntactic_checker/code/services/syntactic_check_service/configuration_handler"
	"syntactic_checker/code/services/syntactic_check_service/results_processor"
)

type syntacticChecksService struct {
	run_configuration  configuration_handler.Configurations
	in_scope_cell_list object_model.ListOfInScopeCells
}

func (syntactic_checks_service *syntacticChecksService) Orchestrate_syntactic_checks(configuration_file_path string) {

	cell_list_checks_service :=
		cell_list_checks_service.
			Create_cell_list_checks_service(
				syntactic_checks_service.in_scope_cell_list,
				syntactic_checks_service.run_configuration.Check_configuration.Issue_types)

	syntactic_check_results :=
		cell_list_checks_service.
			Get_syntactic_check_results()

	results_processor.
		Report_syntactic_check_outputs(
			syntactic_check_results,
			&syntactic_checks_service.run_configuration)
}
