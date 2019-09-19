package identified_string_list_preparers

type IdentifiedStringListPreparerFactory struct{}

func (IdentifiedStringListPreparerFactory) Create(
	csv_filename string,
	check_column_name string,
	identity_column_name string) *identifiedStringListPreparers {

	identified_string_list_preparer :=
		new(
			identifiedStringListPreparers)

	identified_string_list_preparer.
		check_column_name =
		check_column_name

	identified_string_list_preparer.
		identity_column_name =
		identity_column_name

	identified_string_list_preparer.
		csv_filename =
		csv_filename

	return identified_string_list_preparer
}
