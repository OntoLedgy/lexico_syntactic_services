package results_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/configurations"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type resultsProcessors struct {
	output_configuration       configurations.OutputConfigurations
	syntactic_checking_results service_results.IdentifiedStringListChecksResults
	//TODO - should this be wrapped into syntactic_checking_results
}

func (
	results_processor *resultsProcessors) Process_syntactic_check_outputs() map[string][][]string {

	identified_string_list_checks_result :=
		results_processor.
			syntactic_checking_results

	syntactic_check_result_report :=
		prepare_syntactic_checks_results_transactions(
			identified_string_list_checks_result)

	return syntactic_check_result_report

}
