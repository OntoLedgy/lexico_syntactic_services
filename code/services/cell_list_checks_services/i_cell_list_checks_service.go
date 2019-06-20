package cell_list_checks_services

import "syntactic_checker/code/object_model/transactions"

type iCellListChecksService interface {
	Set_syntactic_checks_results()
	Get_fix_transactions() transactions.FixTransactions
	Get_issue_transactions() transactions.IssuesTransactions
}
