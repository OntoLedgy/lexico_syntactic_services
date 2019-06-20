package issues_processor

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
)

//TODO - Stage 2 - convert to type and move data up to struct

type issuesProcessor struct {
	in_scope_issue_types          []issues.IssueTypes
	in_scope_cell                 object_model.Cells
	cell_check_issues_transaction transactions.IssuesTransactions
}

func (
	issues_processor *issuesProcessor) Set_cell_check_issues() {

	in_scope_issue_types :=
		issues_processor.
			in_scope_issue_types

	for _, in_scope_issue_type := range in_scope_issue_types {

		cell_check_issue :=
			issues_processor.
				get_cell_check_issue(
					in_scope_issue_type)

		issues_processor.
			append_cell_check_issue(
				cell_check_issue)
	}

}

func (
	issues_processor *issuesProcessor) Get_cell_check_issue_transactions() transactions.IssuesTransactions {

	return issues_processor.cell_check_issues_transaction

}
