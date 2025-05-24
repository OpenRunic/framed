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

// At retrieves row at n index
func (t *Table) At(n int) *Row {
	return t.Rows[n]
}

// Find retrieves row at using callback
func (t *Table) Find(cb func(*Row) bool) *Row {
	for _, row := range t.Rows {
		if cb(row) {
			return row
		}
	}

	return nil
}

// FindX retrieves row with specific column value
func (t *Table) FindX(column string, value any) *Row {
	idx := t.State.Index(column)

	return t.Find(func(r *Row) bool {
		return r.At(idx) == value
	})
}

// FindAll retrieves multiple rows at using callback
func (t *Table) FindAll(cb func(*Row) bool) []*Row {
	rows := make([]*Row, 0)

	for _, row := range t.Rows {
		if cb(row) {
			rows = append(rows, row)
		}
	}

	return rows
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
