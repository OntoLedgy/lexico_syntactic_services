package contract

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type IStringChecksServices interface {
	Set_string_checks_result()
	Get_string_checks_result() *service_results.StringChecksResults
	Get_string_checks_input() *service_inputs.StringChecksInputs
	Set_string_issues_result(*service_results.IssueChecksResultLists)
	Set_string_fixes_result(service_results.FixChecksResults)
}
