package framed

// pipeline action to cherry-pick columns
type ActionSelection struct {
	name     string
	columns  []string
	callback func(*Table, []string) []string
}

func (a ActionSelection) ExecName() string {
	if len(a.name) > 0 {
		return a.name
	}
	return "selection"
}

func (a ActionSelection) readColumns(src *Table) []string {
	if a.callback != nil {
		return a.callback(src, a.columns)
	}
	return a.columns
}

func (a ActionSelection) Execute(src *Table) (*Table, error) {
	return CherryPick(src, a.readColumns(src), src.Rows)
}

func ColumnSelection(name string, columns ...string) *ActionSelection {
	return &ActionSelection{
		name:    name,
		columns: columns,
	}
}

func ColumnSelectionCallback(name string, callback func(*Table, []string) []string, columns ...string) *ActionSelection {
	return &ActionSelection{name, columns, callback}
}

// pipeline action to pick only provided columns
func PickColumn(columns ...string) *ActionSelection {
	return ColumnSelection("pick_column", columns...)
}

// pipeline action to drop provided columns
func DropColumn(columns ...string) *ActionSelection {
	return ColumnSelectionCallback("drop_column", func(src *Table, s []string) []string {
		return SliceOmit(src.State.Columns, s)
	}, columns...)
}
