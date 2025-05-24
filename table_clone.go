package framed

// Clone duplicates table as-is
func (t *Table) Clone() *Table {
	df := t.CloneE()
	df.Rows = append(df.Rows, t.Rows...)
	return df
}

// CloneE duplicates table as-is without rows
func (t *Table) CloneE() *Table {
	return &Table{
		resolved: t.resolved,
		Rows:     make([]*Row, 0),
		State:    t.State.Clone(),
		Options:  t.Options.Clone(),
	}
}

// CherryPick selects the columns and assigns new
// rows for those columns to build new [Table]
func CherryPick(src *Table, columns []string, rows []*Row) (*Table, error) {
	df := src.CloneE().MarkUnresolved()

	if !src.State.IsEmpty() {
		df.UseColumns(SlicePick(src.State.Columns, columns))

		for _, fRow := range rows {
			df.AddRow(fRow.CloneP(src.State.ResolveIndexes(columns)))
		}
	}

	return df, nil
}
