package framed

type ModifyTableDataReader = func(*State, *Row) *Row

type ActionModifyRow struct {
	callback ModifyTableDataReader
}

func (a ActionModifyRow) ActionName() string {
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

// ModifyRow iterates through all rows to modify the
// row and generates the new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.ModifyRow(func(*State, *Row) *Row),
//		...
//	)
func ModifyRow(cb ModifyTableDataReader) *ActionModifyRow {
	return &ActionModifyRow{
		callback: cb,
	}
}
