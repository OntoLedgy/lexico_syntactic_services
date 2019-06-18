package syntactic_check_service

import (
	"syntactic_checker/code/services/syntactic_check_service/cells_preparer"
	"syntactic_checker/code/services/syntactic_check_service/configuration_handler"
)

func Create_syntactic_check_orchestrator(configuration_file_path string) syntacticChecksService {

	var syntactic_check_orchestrator syntacticChecksService

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
