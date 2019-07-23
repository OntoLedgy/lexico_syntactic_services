package results_processors

import (
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
)

type ResultsProcessorFactory struct{}

func (
	ResultsProcessorFactory) Create(
	cell_list_checks_result service_results.CellListChecksResults,
	output_configuration configuration_getters.OutputConfigurations,
) *resultsProcessors {

	results_processor :=
		new(
			resultsProcessors)

	results_processor.
		output_configuration =
		output_configuration

	results_processor.
		syntactic_checking_results =
		cell_list_checks_result

	return results_processor

}
