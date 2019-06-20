package transactions

import "syntactic_checker/code/object_model/issues"

type SyntacticCheckTransactions struct {
	Fix_transactions   FixTransactions
	Issue_transactions IssuesTransactions
	Issue_parameters   issues.IssueParameters
}
