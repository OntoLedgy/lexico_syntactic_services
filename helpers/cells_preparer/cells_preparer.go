package cells_preparer

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/helpers/configuration_handler"
	"syntactic_checker/object_model"
)

func Get_in_scope_identified_cells(
	run_configuration *configuration_handler.Configurations) object_model.InScopeCells {

	csv_filename :=
		run_configuration.Check_configuration.Csv_file_name

	//TODO - Stage 2 - replace this with arguments passed into this function, function should not know about internals of configuration data structure
	identity_colunmn_name := run_configuration.Check_configuration.Identity_column_name
	cell_value_column_name := run_configuration.Check_configuration.Check_column_name

	in_scope_cells :=
		prepare_in_scope_identified_cells_data(
			csv_filename,
			identity_colunmn_name,
			cell_value_column_name)

	return in_scope_cells
}

func prepare_in_scope_identified_cells_data(
	csv_filename string,
	identity_colunmn_name string,
	cell_value_column_name string) object_model.InScopeCells {

	fmt.Printf(
		"\nReading CSV Data..")

	var in_scope_cells object_model.InScopeCells

	in_scope_identified_cells_raw :=
		storage.Read_csv_data(
			csv_filename)

	fmt.Printf(
		"Preparing extracted data for checks (converting to interface)")

	in_scope_identified_cells_interface :=
		storage_slices.Convert_2d_string_to_interface(
			in_scope_identified_cells_raw)

	in_scope_identified_cells_with_headers :=
		storage.Get_csv_with_headers(
			in_scope_identified_cells_interface)

	//TODO - Stage 2 - convert to Cells data structure

	in_scope_cells.Cells = make([]object_model.InScopeCell, len(in_scope_identified_cells_interface))

	for index, value := range in_scope_identified_cells_with_headers {

		in_scope_cells.Cells[index].Cell_identifier = value[identity_colunmn_name].(string)
		in_scope_cells.Cells[index].Cell_value = value[cell_value_column_name].(string)

	}

	return in_scope_cells
}
