package cell_checks_result_setters

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_checks_services/contract"
	"syntactic_checker/code/services/cell_checks_services/internal/cell_fix_getters"
	"syntactic_checker/code/services/cell_checks_services/internal/cell_issues_getters"
)

type CellChecksResultSetters struct {
	in_scope_cell                cells.Cells
	list_of_in_scope_issue_types []issues.IssueTypes
}

//TODO - convert this to method

func Set_cell_issues_and_fix(
	cell_checks_service contract.ICellChecksServices) {

	cell_checks_parameters := cell_checks_service.Get_cell_checks_parameter()

	in_scope_cell := cell_checks_parameters.In_scope_cell

	list_of_in_scope_issue_types := cell_checks_parameters.List_of_in_scope_issue_types

	cell_checks_issues :=
		cell_issues_getters.
			Get_cell_issues(
				in_scope_cell,
				list_of_in_scope_issue_types)

	there_are_issues :=
		cell_checks_issues != nil

	if there_are_issues {

		cell_checks_service.
			Set_issues_result(
				cell_checks_issues)

		cell_check_fix :=
			get_cell_fix(
				cell_checks_parameters)

		cell_checks_service.
			Set_fixes_result(
				cell_check_fix)

	}
}

func get_cell_fix(cell_checks_parameter service_parameters.CellChecksParameters) fixes.Fixes {

	fix_processor_factory :=
		new(
			cell_fix_getters.FixProcessorsFactory)

	fix_processor :=
		fix_processor_factory.
			Create()

	cell_check_fix :=
		fix_processor.
			Get_cell_check_fix(
				cell_checks_parameter.List_of_in_scope_issue_types,
				cell_checks_parameter.In_scope_cell)

	return cell_check_fix

}
