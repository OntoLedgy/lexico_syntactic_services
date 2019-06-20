package syntactic_check_services

import (
	"syntactic_checker/code/services/syntactic_check_services/cells_preparer"
	"syntactic_checker/code/services/syntactic_check_services/configuration_handler"
)

func Create_syntactic_check_orchestrator(
	configuration_file_path string) ISyntacticChecksService {

	syntactic_check_orchestrator :=
		new(
			syntacticChecksService)

	syntactic_check_orchestrator.
		run_configuration =
		*configuration_handler.
			Get_configuration(
				configuration_file_path)

	syntactic_check_orchestrator.
		in_scope_cell_list =
		cells_preparer.
			Get_in_scope_identified_cells(
				&syntactic_check_orchestrator.run_configuration)

	return syntactic_check_orchestrator
}
