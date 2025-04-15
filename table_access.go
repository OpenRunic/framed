package framed

// Length return number of total rows
func (t *Table) Length() int {
	return len(t.Rows)
}

// ColLength returns number of columns
func (t *Table) ColLength() int {
	return len(t.State.Columns)
}

// IsEmpty verifies if table is empty
func (t *Table) IsEmpty() bool {
	return t.Length() == 0 || t.ColLength() == 0
}

// At retrieves row at x index
func (t *Table) At(idx int) *Row {
	return t.Rows[idx]
}

// IsAtMaxLine checks rows are already at restricted max limit
func (t *Table) IsAtMaxLine() bool {
	if t.Options.MaxRows < 0 {
		return false
	}

	return t.Length() >= t.Options.MaxRows
}

// IsResolved checks table is resolved
func (t *Table) IsResolved() bool {
	return t.resolved
}
