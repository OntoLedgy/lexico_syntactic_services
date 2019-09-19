package contract

import (
	"syntactic_checker/code/object_model/configurations"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type ISyntacticCheckingServices interface {
	Run_syntactic_checking_service()
	Get_run_configuration() *configurations.RunConfigurations
	Set_syntactic_check_results(syntactic_checking_result service_results.IdentifiedStringListChecksResults)
	Get_identified_string_list() *identified_strings.IdentifiedStringLists
	Get_syntactic_checking_result() service_results.IdentifiedStringListChecksResults
}
