package syntactic_check_services

import (
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/syntactic_check_services/configuration_handler"
)

type ISyntacticChecksService interface {
	Orchestrate_syntactic_checks(configuration_file_path string)
	Get_syntactic_check_transactions() transactions.SyntacticCheckTransactions
	Get_syntactic_check_run_configuration() configuration_handler.Configurations
}
