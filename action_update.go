package framed

// pipeline action to update options and details on table
type ActionUpdate struct {
	callback func(*Table) (*Table, error)
}

func (a ActionUpdate) Execute(src *Table) (*Table, error) {
	return a.callback(src)
}

func ChangeOptions(cb func(*Table) (*Table, error)) *ActionUpdate {
	return &ActionUpdate{cb}
}
