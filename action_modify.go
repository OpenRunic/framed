package framed

type ModifyTableDataReader = func(*State, *Row) *Row

// pipeline action to modify every row in table
type ActionModifyRow struct {
	callback ModifyTableDataReader
}

func (a ActionModifyRow) ExecName() string {
	return "modify_row"
}

func (a ActionModifyRow) Execute(src *Table) (*Table, error) {
	df := src.CloneE()

	for _, r := range src.Rows {
		row := r.Clone()
		df.AddRow(a.callback(src.State, row))
	}

	return df, nil
}

func ModifyRow(cb ModifyTableDataReader) *ActionModifyRow {
	return &ActionModifyRow{
		callback: cb,
	}
}
