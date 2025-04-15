package framed

type ActionSelection struct {
	name     string
	columns  []string
	callback func(*Table, []string) []string
}

func (a ActionSelection) ActionName() string {
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

// PickColumn plucks the specified columns and
// generates the new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.PickColumn("col1", "col2", ...),
//		...
//	)
func PickColumn(columns ...string) *ActionSelection {
	return ColumnSelection("pick_column", columns...)
}

// DropColumn ignores the specified columns and
// generates the new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.DropColumn("col1", "col2", ...),
//		...
//	)
func DropColumn(columns ...string) *ActionSelection {
	return ColumnSelectionCallback("drop_column", func(src *Table, s []string) []string {
		return SliceOmit(src.State.Columns, s)
	}, columns...)
}
