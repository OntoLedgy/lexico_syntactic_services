package result_setters

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/identified_string_list_checks_services"
	"syntactic_checker/code/services/syntactic_checking_services/contract"
)

//move to separate package

type SyntacticCheckingServicesResultSetters struct {
	contract.ISyntacticCheckingServices
}

func (
	syntactic_checking_service_setter *SyntacticCheckingServicesResultSetters) Set_syntactic_checking_result() {

	identfied_string_list_checks_result :=
		syntactic_checking_service_setter.
			get_identified_string_list_checks_result()

	syntactic_checking_service_setter.
		Set_syntactic_check_results(
			identfied_string_list_checks_result)

}

func (
	syntactic_checking_service_setter *SyntacticCheckingServicesResultSetters) get_identified_string_list_checks_result() service_results.IdentifiedStringListChecksResults {

	identified_string_list_checks_parameter :=
		syntactic_checking_service_setter.
			generate_identified_string_list_checks_service_parameter()

	identified_string_list_checks_service_factory :=
		new(
			identified_string_list_checks_services.
				IdentifiedStringListChecksServiceFactory)

	identified_string_list_checks_service :=
		identified_string_list_checks_service_factory.
			Create(
				*identified_string_list_checks_parameter)

	identified_string_list_checks_service.
		Set_syntactic_checks_results()

	identified_string_list_checks_result :=
		identified_string_list_checks_service.
			Get_identified_string_list_checks_result()

	return identified_string_list_checks_result
}

func (
	syntactic_checking_service_setter *SyntacticCheckingServicesResultSetters) generate_identified_string_list_checks_service_parameter() *service_parameters.IdentifiedStringListChecksParameters {

	identified_string_list :=
		syntactic_checking_service_setter.
			Get_identified_string_list()

	run_configuration :=
		syntactic_checking_service_setter.
			Get_run_configuration()

	in_scope_issue_types :=
		run_configuration.
			Check_configuration.
			Issue_types

	identified_string_list_checks_parameter :=
		new(
			service_parameters.
				IdentifiedStringListChecksParameters)

	identified_string_list_checks_parameter.
		Identified_string_list =
		identified_string_list

	identified_string_list_checks_parameter.
		List_of_in_scope_issue_types =
		in_scope_issue_types

	return identified_string_list_checks_parameter
}
