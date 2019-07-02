package cells_preparers

type CellsPreparerFactory struct{}

func (CellsPreparerFactory) Create(
	csv_filename string,
	check_column_name string,
	identity_column_name string) *cellsPreparers {

	cell_preparer := new(
		cellsPreparers)

	cell_preparer.
		check_column_name =
		check_column_name

	cell_preparer.
		identity_column_name =
		identity_column_name

	cell_preparer.
		csv_filename =
		csv_filename

	return cell_preparer
}
