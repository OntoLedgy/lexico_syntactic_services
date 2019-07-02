package cell_checks_services

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_checks_services/internal/fix_processors"
	"syntactic_checker/code/services/cell_checks_services/internal/issues_processor"
)

func (
	cell_checks_service *cellChecksService) check_cell_for_issues() []issues.Issues {

	issues_processor :=
		issues_processor.
			Create(
				cell_checks_service.Cell_checks_parameter.In_scope_cell,
				cell_checks_service.Cell_checks_parameter.List_of_in_scope_issue_types)

	issues_processor.
		Set_cell_check_issues()

	cell_checks_issues :=
		issues_processor.Get_cell_checks_issues()

	return cell_checks_issues
}

func (
	cell_checks_service *cellChecksService) set_cell_issues_and_fix(
	cell_checks_issues []issues.Issues) {

	there_are_issues :=
		cell_checks_issues != nil

	if there_are_issues {

		cell_checks_service.
			Cell_checks_result.
			Cell_checks_issues =
			cell_checks_issues

		cell_checks_service.
			set_cell_fix()

	}
}

func (
	cell_checks_service *cellChecksService) set_cell_fix() {

	fix_processor_factory := new(fix_processors.FixProcessorsFactory)

	fix_processor :=
		fix_processor_factory.Create()

	cell_check_fix :=
		fix_processor.
			Get_cell_check_fix(
				cell_checks_service.Cell_checks_parameter.List_of_in_scope_issue_types,
				cell_checks_service.Cell_checks_parameter.In_scope_cell)

	cell_checks_service.
		Cell_checks_result.
		Cell_checks_fix =
		cell_check_fix

}
