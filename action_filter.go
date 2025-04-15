package framed

type ActionFilterRow struct {
	callback func(*State, *Row) bool
}

func (a ActionFilterRow) ActionName() string {
	return "filter"
}

func (a ActionFilterRow) Execute(src *Table) (*Table, error) {
	if src.IsEmpty() {
		return src.CloneE(), nil
	}

	idx := 0
	var filtered = make([]*Row, 0)
	for _, row := range src.Rows {
		if a.callback(src.State, row) {
			filtered = append(filtered, row.Clone().WithIndex(idx))
			idx++
		}
	}

	return CherryPick(src, src.State.Columns, filtered)
}

// FilterRow iterates through all rows, filters the rows
// and build a new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.FilterRow(func(*framed.State, *framed.Row) bool),
//		...
//	)
func FilterRow(cb func(*State, *Row) bool) *ActionFilterRow {
	return &ActionFilterRow{cb}
}
