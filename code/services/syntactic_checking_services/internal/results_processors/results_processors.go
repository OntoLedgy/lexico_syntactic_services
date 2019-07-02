package results_processors

import (
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/configuration_getters"
)

type resultsProcessors struct {
	output_configuration       configuration_getters.OutputConfigurations
	syntactic_checking_results service_results.CellListChecksResults
	//should this be wrapped into syntactic_checking_results
}
