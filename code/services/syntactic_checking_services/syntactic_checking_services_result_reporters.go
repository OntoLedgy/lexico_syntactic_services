package syntactic_checking_services

import "syntactic_checker/code/services/syntactic_checking_services/internal/results_processors"

type syntacticCheckingResultReporters struct {
	syntacticCheckingServices
}

func (
	syntactic_checking_results_reporter *syntacticCheckingResultReporters) report_syntactic_checking_result() {

	output_configuration :=
		syntactic_checking_results_reporter.
			run_configuration.
			Output_configuration

	cell_list_checks_result :=
		syntactic_checking_results_reporter.
			syntactic_checking_result

	results_processor_factory :=
		new(
			results_processors.
				ResultsProcessorFactory)

	results_processor :=
		results_processor_factory.
			Create(
				cell_list_checks_result,
				output_configuration)

	results_processor.
		Report_syntactic_check_outputs()
}
