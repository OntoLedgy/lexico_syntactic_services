package result_reporters

import (
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators/result_reporters/results_processors"
)

type SyntacticCheckingResultReporters struct {
	contract.ISyntacticCheckingServices
}

func (
	syntactic_checking_results_reporter *SyntacticCheckingResultReporters) Report_syntactic_checking_result() {

	run_configuration :=
		syntactic_checking_results_reporter.
			Get_run_configuration()

	output_configuration :=
		run_configuration.
			Output_configuration

	cell_list_checks_result :=
		syntactic_checking_results_reporter.
			Get_syntactic_checking_result()

	results_processor_factory :=
		new(
			results_processors.
				ResultsProcessorFactory)

	results_processor :=
		results_processor_factory.
			Create(
				cell_list_checks_result,
				output_configuration)

	syntactic_check_result_report :=
		results_processor.
			Process_syntactic_check_outputs()

	syntactic_checking_results_reporter.
		publish_syntactic_checking_result(
			syntactic_check_result_report)
}

func (
	syntactic_checking_results_reporter *SyntacticCheckingResultReporters) publish_syntactic_checking_result(
	syntactic_check_result_report map[string][][]string) {

	output_configuration :=
		syntactic_checking_results_reporter.
			Get_run_configuration().
			Output_configuration

	issues_file_name :=
		output_configuration.
			Output_issues_file_absolute_path

	issue_parameters_file_name :=
		output_configuration.
			Output_issue_parameters_file_absolute_path

	fixes_file_name :=
		output_configuration.
			Output_fixes_file_absolute_path

	write_syntactic_checking_result_to_csvs(
		syntactic_check_result_report,
		issues_file_name,
		issue_parameters_file_name,
		fixes_file_name)
}
