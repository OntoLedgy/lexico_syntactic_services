package service_run_preparers

import (
	"github.com/OntoLedgy/logging_services/standard_global_logger"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/configurations"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/services/service_run_preparers/internal/configuration_getters"
	"github.com/OntoLedgy/syntactic_checker/code/services/service_run_preparers/internal/configuration_getters/identified_string_list_preparers"
)

type ServiceRunPreparers struct {
	service_type            string
	configuration_file_path string
}

func (
	service_run_preparer *ServiceRunPreparers) Get_service_run_data(
	configuration_file_path string) *configurations.SyntacticCheckingData {

	syntactic_checking_data :=
		new(
			configurations.SyntacticCheckingData)

	run_configuration :=
		service_run_preparer.
			get_current_run_configuration(
				configuration_file_path)

	standard_global_logger.
		Start_logger(run_configuration.Output_configuration.Output_log_folder_absolute_path,
			"go_syntactic_checker")

	identified_string_list :=
		service_run_preparer.
			load_identified_string_list(
				run_configuration.
					Check_configuration)

	syntactic_checking_data.
		Run_configuration =
		&run_configuration

	syntactic_checking_data.
		Identified_string_list =
		&identified_string_list

	return syntactic_checking_data

}

func (
	ServiceRunPreparers) get_current_run_configuration(
	configuration_file_path string) configurations.RunConfigurations {

	configuration_getter_factory :=
		new(
			configuration_getters.
				ConfigurationGetterFactories)

	configuration_getter :=
		configuration_getter_factory.
			Create()

	run_configuration :=
		*configuration_getter.
			Get_configuration(
				configuration_file_path)

	return run_configuration
}

func (
	ServiceRunPreparers) load_identified_string_list(
	check_configuration configurations.CheckConfigurations) identified_strings.IdentifiedStringLists {

	identity_column_name :=
		check_configuration.
			Identity_column_name

	check_column_name :=
		check_configuration.
			Check_column_name

	csv_filename :=
		check_configuration.
			Input_csv_file_name

	identified_string_list_preparer_factory :=
		new(
			identified_string_list_preparers.
				IdentifiedStringListPreparerFactory)

	identified_string_list_preparer :=
		identified_string_list_preparer_factory.
			Create(
				csv_filename,
				check_column_name,
				identity_column_name)

	return identified_string_list_preparer.Get_in_scope_identified_string_list()

}
