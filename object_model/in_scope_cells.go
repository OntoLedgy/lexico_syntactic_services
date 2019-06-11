package object_model

type InScopeCells struct {
	Cells []InScopeCell
}

type InScopeCell struct {
	Cell_identifier string
	Cell_value      string
}
