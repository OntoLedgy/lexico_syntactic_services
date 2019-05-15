package checks_runner

import (
	"database_manager/utils"
	"fmt"
	"syntactic_checker/cell_checks_orchestrator/cell_fixer"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
	"syntactic_checker/cell_checks_orchestrator/regex_processor"
	"syntactic_checker/object_model"
)

func process_syntactic_checks_for_cells(
	in_scope_identified_cells []map[string]interface{},
	run_configuration *configuration_handler.Configurations) ([][]interface{}, [][]interface{}) {

	var cells_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_fix_transaction []interface{}
	var cells_syntactic_check_fix_transactions [][]interface{}

	fmt.Printf(
		"processing checks: %s\n",
		run_configuration)

	for _, in_scope_cell := range in_scope_identified_cells {

		cell_syntactic_check_issues_transactions, cell_syntactic_check_fix_transaction =
			process_syntactic_checks_for_cell(
				in_scope_cell,
				run_configuration)

		cells_syntactic_check_issues_transactions =
			append(
				cells_syntactic_check_issues_transactions,
				cell_syntactic_check_issues_transactions...)

		if cell_syntactic_check_fix_transaction != nil {

			cells_syntactic_check_fix_transactions =
				append(
					cells_syntactic_check_fix_transactions,
					cell_syntactic_check_fix_transaction)

			cell_syntactic_check_fix_transaction = nil

		}

	}

	/*
		if cells_syntactic_check_issues_transactions != nil { //#TODO add to logger
			fmt.Printf(
				"\n++++++++++++++++++++++++Issue_types completed: %s issues found.\nExample transaction row: %s ",
				len(cells_syntactic_check_issues_transactions),
				cells_syntactic_check_issues_transactions[0])
		}*/
	return cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions
}

func process_syntactic_checks_for_cell(
	in_scope_identified_cell map[string]interface{},
	run_configuration *configuration_handler.Configurations) ([][]interface{}, []interface{}) {

	var cell_syntactic_check_issue_transactions [][]interface{}
	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	var cell_syntactic_check_issue_results []object_model.Regex_check_results

	in_scope_syntactic_check_types :=
		run_configuration.Check_configuration.Issue_types

	cell_syntactic_check_issue_transactions, cell_syntactic_check_issue_results =
		process_syntactic_check_issues_for_cell(
			in_scope_syntactic_check_types,
			cell_syntactic_check_issue_transactions,
			in_scope_identified_cell,
			run_configuration.Check_configuration.Check_column_name,
			run_configuration.Check_configuration.Identity_column_name)

	cell_syntactic_check_aggregated_fixes_transaction =
		process_syntactic_check_fixes_transaction_for_cell(
			in_scope_syntactic_check_types,
			in_scope_identified_cell,
			run_configuration.Check_configuration.Check_column_name,
			run_configuration.Check_configuration.Identity_column_name,
			cell_syntactic_check_issue_results)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction
}

func process_syntactic_check_issues_for_cell(
	in_scope_syntactic_check_types []object_model.Issue_types,
	cell_syntactic_check_issue_transactions [][]interface{},
	in_scope_identified_cell map[string]interface{},
	check_column_name string,
	identity_column_name string) ([][]interface{}, []object_model.Regex_check_results) {

	var cell_syntactic_check_issue_results []object_model.Regex_check_results

	for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

		cell_syntactic_check_issue_result :=
			process_syntactic_check_issue_for_cell(
				in_scope_identified_cell,
				check_column_name,
				in_scope_syntactic_check_type)

		cell_syntactic_check_issue_results, cell_syntactic_check_issue_transactions =
			aggregate_syntactic_check_issues(
				in_scope_identified_cell,
				identity_column_name,
				in_scope_syntactic_check_type,
				cell_syntactic_check_issue_results,
				cell_syntactic_check_issue_result,
				cell_syntactic_check_issue_transactions)
	}

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_issue_results

}

func process_syntactic_check_issue_for_cell(
	in_scope_identified_cell map[string]interface{},
	check_column_name string,
	in_scope_check object_model.Issue_types) *object_model.Regex_check_results {

	var regex_check_result *object_model.Regex_check_results

	if in_scope_identified_cell[check_column_name] != nil {

		regex_check_result = //TODO - Stage 3 - add switch to include non-regex in_scope_check types.
			regex_processor.
				Process_regex_check( //TODO - Stage 2 - replace the check interface with the check type object and pass it through
					in_scope_check.Issue_check_regex,
					in_scope_identified_cell[check_column_name], // cell value
					in_scope_check.Issue_check_replacement_string)

	} else {
		fmt.Printf(
			"\nWARNING: cell value for row_id%s is null\n",
			in_scope_identified_cell[check_column_name])

	}
	return regex_check_result
}

func process_syntactic_check_fixes_transaction_for_cell(
	in_scope_syntactic_check_types []object_model.Issue_types,
	in_scope_identified_cell map[string]interface{},
	check_column_name string,
	identity_column_name string,
	cell_syntactic_check_issue_results []object_model.Regex_check_results) []interface{} {

	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	var interim_cell_value_modified, interim_cell_value_marked map[string]interface{}

	if cell_syntactic_check_issue_results != nil {

		fmt.Printf(
			"\nprocessing fixes...\n")

		fix_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		interim_cell_value_modified = CopyMap(in_scope_identified_cell)
		interim_cell_value_marked = CopyMap(in_scope_identified_cell)

		for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

			cell_syntactic_check_issue_result :=
				process_syntactic_check_issue_for_cell(
					interim_cell_value_modified,
					check_column_name,
					in_scope_syntactic_check_type)

			if cell_syntactic_check_issue_result != nil {

				interim_cell_value_modified[check_column_name] =
					cell_fixer.Modify_string_by_index(
						interim_cell_value_modified[check_column_name].(string),
						cell_syntactic_check_issue_result.Replacement_string,
						cell_syntactic_check_issue_result.Regex_match_indices)

				interim_cell_value_marked[check_column_name] =
					cell_fixer.Modify_string_by_index(
						interim_cell_value_marked[check_column_name].(string),
						cell_syntactic_check_issue_result.Mark_string,
						cell_syntactic_check_issue_result.Regex_match_indices)

			}

		}

		cell_syntactic_check_aggregated_fixes_transaction =
			append(
				cell_syntactic_check_aggregated_fixes_transaction,
				in_scope_identified_cell[check_column_name],
				interim_cell_value_marked[check_column_name],
				interim_cell_value_modified[check_column_name])

		cell_syntactic_check_aggregated_fixes_transaction =
			append(
				cell_syntactic_check_aggregated_fixes_transaction,
				fix_uuid.String(),
				in_scope_identified_cell[identity_column_name])

		fmt.Printf("\ncell fix transaction: %v \n", cell_syntactic_check_aggregated_fixes_transaction)

	}

	return cell_syntactic_check_aggregated_fixes_transaction
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
