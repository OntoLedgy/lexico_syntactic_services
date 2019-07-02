package syntactic_checking_services

import (
	"syntactic_checker/code/services/syntactic_checking_services/configuration_getters"
	"syntactic_checker/code/services/syntactic_checking_services/internal/cells_preparers"
)

type SyntacticCheckingServiceFactory struct{}

func (
	factory SyntacticCheckingServiceFactory) Create(
	configuration_file_path string) ISyntacticCheckingService {

	syntactic_checking_service :=
		new(
			syntacticCheckingServices)

	configuration_getter_factory :=
		new(
			configuration_getters.
				ConfigurationGetterFactories)

	configuration_getter :=
		configuration_getter_factory.
			Create()

	syntactic_checking_service.
		run_configuration =
		*configuration_getter.
			Get_configuration(
				configuration_file_path)

	factory.
		load_in_scope_cell_list(
			syntactic_checking_service)

	return syntactic_checking_service
}

func (
	SyntacticCheckingServiceFactory) load_in_scope_cell_list(
	syntactic_checking_service *syntacticCheckingServices) {

	check_configuration :=
		syntactic_checking_service.
			run_configuration.
			Check_configuration

	identity_column_name :=
		check_configuration.
			Identity_column_name

	check_column_name :=
		check_configuration.
			Check_column_name

	csv_filename :=
		check_configuration.
			Input_csv_file_name

	cells_preparer_factory :=
		new(
			cells_preparers.CellsPreparerFactory)

	cells_preparer :=
		cells_preparer_factory.
			Create(
				csv_filename,
				check_column_name,
				identity_column_name)

	syntactic_checking_service.
		in_scope_cell_list =
		cells_preparer.
			Get_in_scope_identified_cells()

}
