package configurations

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type SyntacticCheckingData struct {
	Run_configuration          *RunConfigurations
	Identified_string_list     *identified_strings.IdentifiedStringLists
	Syntactic_checking_results *service_results.IdentifiedStringListChecksResults
}
