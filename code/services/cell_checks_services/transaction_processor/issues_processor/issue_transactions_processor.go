package issues_processor

import (
	"fmt"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services"
)

func process_issue_transactions(
	cell_check_service *cell_check_services.CellCheckService) *issues.Issues {

	if cell_check_service.Regex_check_result != nil {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_check_issue :=
			generate_issue_transaction(
				cell_check_service)

		return &cell_check_issue
	}

	return nil
}

func generate_issue_transaction(
	cell_check_service *cell_check_services.CellCheckService) issues.Issues {

	var cell_check_issue issues.Issues

	cell_check_issue.
		Object_uuid =
		cell_check_issue.
			Set_object_uuid()

	cell_check_issue.
		Issue_type =
		cell_check_service.
			Issue_type

	cell_check_issue.
		Cell =
		cell_check_service.
			In_scope_cell

	cell_check_issue.
		Modified_cell_value =
		cell_check_service.
			Cell_value_edit_history.
			Get_modified_string()

	cell_check_issue.
		Marked_cell_value =
		cell_check_service.
			Cell_value_edit_history.
			Get_marked_string()

	return cell_check_issue
}
