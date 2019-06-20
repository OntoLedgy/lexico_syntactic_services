package results_processor

import (
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/syntactic_check_services/configuration_handler"
)

type resultsProcessor struct {
	syntactic_check_transactions transactions.SyntacticCheckTransactions
	output_configuration         configuration_handler.Output_configurations
}
