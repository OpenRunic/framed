package framed

type ActionUpdate struct {
	callback func(*Table) (*Table, error)
}

func (a ActionUpdate) ActionName() string {
	return "update_table"
}

func (a ActionUpdate) Execute(src *Table) (*Table, error) {
	return a.callback(src)
}

// UpdateTable provides access to current table in pipeline
// and allows to change options and definitions as required.
//
//	newTable, err := table.Execute(
//		...
//		framed.UpdateTable(func(*Table) (*Table, error)),
//		...
//	)
func UpdateTable(cb func(*Table) (*Table, error)) *ActionUpdate {
	return &ActionUpdate{cb}
}
