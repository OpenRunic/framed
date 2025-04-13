package framed

// clone table as-is
func (t *Table) Clone() *Table {
	df := t.CloneE()
	df.Rows = append(df.Rows, t.Rows...)
	return df
}

// clone table with state and options but resolved
func (t *Table) CloneO() *Table {
	return New().
		SetState(t.State.Clone()).
		SetOptions(t.Options.Clone())
}

// clone table as-is without rows
func (t *Table) CloneE() *Table {
	return &Table{
		resolved: t.resolved,
		Rows:     make([]*Row, 0),
		State:    t.State.Clone(),
		Options:  t.Options.Clone(),
	}
}

// cherry pick the columns and build table
func CherryPick(src *Table, columns []string, rows []*Row) (*Table, error) {
	df := src.CloneO()

	if !src.State.IsEmpty() {
		df.UseColumns(SlicePick(src.State.Columns, columns))

		for _, fRow := range rows {
			nRow, err := fRow.CloneP(src.State, columns...)
			if err != nil {
				return nil, err
			}

			df.AddRow(nRow)
		}
	}

	return df, nil
}
