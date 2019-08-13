package contract

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type IIdentifiedStringListChecksServices interface {
	Set_syntactic_checks_results()
	Get_identified_string_list_checks_result() service_results.IdentifiedStringListChecksResults
	Set_identified_string_checks_result(identified_strings.IdentifiedStrings, service_results.StringChecksResults)
	Get_identified_string_list_checks_parameter() service_parameters.IdentifiedStringListChecksParameters
}
