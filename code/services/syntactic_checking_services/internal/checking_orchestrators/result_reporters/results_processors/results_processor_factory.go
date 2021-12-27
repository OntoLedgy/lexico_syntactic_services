package results_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/configurations"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type ResultsProcessorFactory struct{}

func (
	ResultsProcessorFactory) Create(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults,
	output_configuration configurations.OutputConfigurations,
) *resultsProcessors {

	results_processor :=
		new(
			resultsProcessors)

	results_processor.
		output_configuration =
		output_configuration

	results_processor.
		syntactic_checking_results =
		identified_string_list_checks_result

	return results_processor

}
