package contract

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/object_model"
)

type ISyntacticCheckingServices interface {
	Run_syntactic_checking_service()
	Get_run_configuration() object_model.RunConfigurations
	Set_syntactic_check_results(syntactic_checking_result service_results.IdentifiedStringListChecksResults)
	Get_identified_string_list() identified_strings.IdentifiedStringLists
	Get_syntactic_checking_result() service_results.IdentifiedStringListChecksResults
}
