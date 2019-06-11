package column_checks

type IColumnCheckers interface {
	RunColumnChecks() map[string][][]string
}
