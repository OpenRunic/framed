package framed

// number of rows
func (t *Table) Length() int {
	return len(t.Rows)
}

// number of columns
func (t *Table) ColLength() int {
	return len(t.State.Columns)
}

// check table is empty
func (t *Table) Empty() bool {
	return t.Length() == 0 || t.ColLength() == 0
}

// get row at provided index
func (t *Table) At(idx int) *Row {
	return t.Rows[idx]
}

// check table is already at restricted max line count
func (t *Table) IsAtMaxLine() bool {
	if t.Options.MaxRows < 0 {
		return false
	}

	return t.Length() >= t.Options.MaxRows
}

// check is resolved
func (t *Table) IsResolved() bool {
	return t.resolved
}
