package cell_checks_runner

import (
	"fmt"
	"storage/slices"
	"syntactic_checker/cell_checks_results_finaliser"
	"syntactic_checker/configuration_handler"
)

func Run(
	in_scope_identified_cells [][]interface{},
	run_configuration *configuration_handler.Configurations) map[string][][]string {

	var syntactic_check_result_report map[string][][]string

	column_uuid :=
		run_configuration.Csv_configuration.Check_column_uuid
	in_scope_checks :=
		get_in_scope_syntactic_checks(run_configuration)

	fmt.Printf(
		"\nStarting processing. \ncolumn: %v\nin_scope_checks: %s\n",
		column_uuid,
		storage.Pretty_print(in_scope_checks))

	syntactic_check_result_transactions :=
		process_syntactic_checks_for_cells(
			in_scope_identified_cells,
			in_scope_checks)

	syntactic_check_result_report =
		cell_checks_results_finaliser.Prepare_syntactic_checks_results_transactions(
			syntactic_check_result_transactions,
			column_uuid)

	return syntactic_check_result_report
}

func get_in_scope_syntactic_checks(
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	var in_scope_check_interface []interface{}
	var in_scope_checks_interface [][]interface{}

	in_scope_checks :=
		run_configuration.Csv_configuration.Issue_types

	for _, in_scope_check := range in_scope_checks {

		in_scope_check_interface = append(in_scope_check_interface,
			in_scope_check.Issue_type_uuid,
			in_scope_check.Issue_type_name,
			in_scope_check.Issue_check_regex,
			in_scope_check.Issue_check_replacement_string)
		in_scope_checks_interface = append(in_scope_checks_interface, in_scope_check_interface)
		in_scope_check_interface = nil
	}

	return in_scope_checks_interface
}
