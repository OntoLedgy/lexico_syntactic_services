package syntactic_check_services

import (
	"syntactic_checker/code/services/cell_list_checks_services"
	"syntactic_checker/code/services/syntactic_check_services/results_processor"
)

func (
	syntactic_checks_service *syntacticChecksService) Orchestrate_syntactic_checks(configuration_file_path string) {

	syntactic_checks_service.
		set_syntactic_check_transactions()

	syntactic_checks_service.
		report_syntactic_check_results()

}

func (
	syntactic_checks_service *syntacticChecksService) set_syntactic_check_transactions() {

	in_scope_cell_list :=
		syntactic_checks_service.
			in_scope_cell_list

	in_scope_issue_types :=
		syntactic_checks_service.
			run_configuration.
			Check_configuration.
			Issue_types

	cell_list_checks_service :=
		cell_list_checks_services.
			Create_cell_list_checks_service(
				in_scope_cell_list,
				in_scope_issue_types)

	cell_list_checks_service.
		Set_syntactic_checks_results()

	syntactic_checks_service.
		Syntactic_check_results_transactions.
		Fix_transactions =
		cell_list_checks_service.
			Get_fix_transactions()

	syntactic_checks_service.
		Syntactic_check_results_transactions.
		Issue_transactions =
		cell_list_checks_service.
			Get_issue_transactions()
}

func (
	syntactic_checks_service *syntacticChecksService) report_syntactic_check_results() {
	results_processor :=
		results_processor.
			Create_results_processor(
				syntactic_checks_service.
					Syntactic_check_results_transactions,
				syntactic_checks_service.
					run_configuration.
					Output_configuration)
	results_processor.
		Report_syntactic_check_outputs()
}
