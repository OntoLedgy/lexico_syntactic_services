package cell_checks_services

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/transactions"
)

type ICellCheckOrchestrators interface {
	Set_cell_checks_results()
	Get_cell_checks_issues() transactions.IssuesTransactions
	Get_cell_checks_fix() fixes.Fixes
}
