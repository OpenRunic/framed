package framed

type ActionUpdate struct {
	callback func(*Table) (*Table, error)
}

func (a ActionUpdate) Execute(src *Table) (*Table, error) {
	return a.callback(src)
}

// ChangeOptions provides access to current table in pipeline
// and allows to change options and definitions as required.
//
//	newTable, err := table.Execute(
//		...
//		framed.ChangeOptions(func(*Table) (*Table, error)),
//		...
//	)
func ChangeOptions(cb func(*Table) (*Table, error)) *ActionUpdate {
	return &ActionUpdate{cb}
}
