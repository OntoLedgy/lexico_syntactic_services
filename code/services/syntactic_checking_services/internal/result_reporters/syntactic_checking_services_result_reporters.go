package result_reporters

import (
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal/results_processors"
)

//move to separate package

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

	results_processor.
		Report_syntactic_check_outputs()
}
