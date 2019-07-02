package cells_preparers

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/code/object_model/cells"
)

type cellsPreparers struct {
	identity_column_name string
	check_column_name    string
	csv_filename         string //should be path?
}

func (
	cell_preparer cellsPreparers) Get_in_scope_identified_cells() cells.ListOfCells {

	csv_filename :=
		cell_preparer.
			csv_filename

	identity_colunmn_name :=
		cell_preparer.
			identity_column_name

	cell_value_column_name :=
		cell_preparer.
			check_column_name

	in_scope_cells :=
		prepare_in_scope_identified_cells_data(
			csv_filename,
			identity_colunmn_name,
			cell_value_column_name)

	return in_scope_cells
}

//TODO - Stage 3 - cell data reading needs to be improved

func prepare_in_scope_identified_cells_data(
	csv_filename string,
	identity_colunmn_name string,
	cell_value_column_name string) cells.ListOfCells {

	fmt.Printf(
		"\nReading CSV Data..")

	var in_scope_cells cells.ListOfCells

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

	in_scope_cells.Cells =
		make(
			[]cells.Cells,
			len(in_scope_identified_cells_interface))

	for index, value := range in_scope_identified_cells_with_headers {

		in_scope_cells.Cells[index].Cell_identifier =
			value[identity_colunmn_name].(string)

		in_scope_cells.Cells[index].Cell_value =
			value[cell_value_column_name].(string)

	}

	return in_scope_cells
}
