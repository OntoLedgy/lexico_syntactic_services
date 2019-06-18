package cell_checks_service

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/transactions"
)

type ICellCheckOrchestrators interface {
	Get_cell_checks_results() ([]interface{}, fixes.Fixes, transactions.IssueTransactions)
}
