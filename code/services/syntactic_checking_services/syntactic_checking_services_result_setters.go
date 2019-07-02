package syntactic_checking_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/cell_list_checks_services"
)

type syntacticCheckingServicesResultSetters struct {
	syntacticCheckingServices
}

func (
	syntactic_checking_service_setter *syntacticCheckingServicesResultSetters) set_syntactic_checking_result() {

	cell_list_checks_result :=
		syntactic_checking_service_setter.
			get_cell_list_checks_result()

	syntactic_checking_service_setter.
		syntactic_checking_result =
		cell_list_checks_result
}

func (
	syntactic_checking_service *syntacticCheckingServices) get_cell_list_checks_result() service_results.CellListChecksResults {

	cell_list_checks_parameter :=
		syntactic_checking_service.
			generate_cell_list_checks_service_parameter()

	cell_list_checks_service_factory :=
		new(
			cell_list_checks_services.
				CellListChecksServiceFactory)

	cell_list_checks_service :=
		cell_list_checks_service_factory.
			Create(
				*cell_list_checks_parameter)

	cell_list_checks_service.
		Set_syntactic_checks_results()

	cell_list_checks_result :=
		cell_list_checks_service.
			Get_cell_list_checks_result()

	return cell_list_checks_result
}

func (
	syntactic_checking_service *syntacticCheckingServices) generate_cell_list_checks_service_parameter() *service_parameters.CellListChecksParameters {

	in_scope_cell_list :=
		syntactic_checking_service.
			in_scope_cell_list

	in_scope_issue_types :=
		syntactic_checking_service.
			run_configuration.
			Check_configuration.
			Issue_types

	cell_list_checks_parameter :=
		new(
			service_parameters.
				CellListChecksParameters)

	cell_list_checks_parameter.
		List_of_in_scope_cells =
		in_scope_cell_list

	cell_list_checks_parameter.
		List_of_in_scope_issue_types =
		in_scope_issue_types

	return cell_list_checks_parameter
}
