package framed

// pipeline action to filter the rows based on callback
type ActionFilter struct {
	callback func(*State, *Row) bool
}

func (a ActionFilter) ExecName() string {
	return "filter"
}

func (a ActionFilter) Execute(src *Table) (*Table, error) {
	if src.Empty() {
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

func Filter(cb func(*State, *Row) bool) *ActionFilter {
	return &ActionFilter{cb}
}
