package results_processors

import (
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
)

type resultsProcessors struct {
	output_configuration       configuration_getters.OutputConfigurations
	syntactic_checking_results service_results.CellListChecksResults
	//should this be wrapped into syntactic_checking_results
}

func (
	results_processor *resultsProcessors) Report_syntactic_check_outputs() {

	//TODO - Split this method

	cell_list_checks_result :=
		results_processor.
			syntactic_checking_results

	syntactic_check_result_report :=
		prepare_syntactic_checks_results_transactions(
			cell_list_checks_result)

	output_configuration :=
		results_processor.output_configuration

	issues_file_name :=
		output_configuration.
			Output_issues_file_absolute_path

	issue_parameters_file_name :=
		output_configuration.
			Output_issue_parameters_file_absolute_path

	fixes_file_name :=
		output_configuration.
			Output_fixes_file_absolute_path

	syntactic_check_issues_set :=
		syntactic_check_result_report["syntactic_check_issues_set"]

	report_syntactic_check_issues(
		syntactic_check_issues_set,
		issues_file_name)

	report_syntactic_check_issue_parameters(
		syntactic_check_result_report["syntactic_check_issue_parameters_set"],
		issue_parameters_file_name)

	report_syntactic_check_fixes(
		syntactic_check_result_report["syntactic_check_fix_transactions_set"],
		fixes_file_name)
}
