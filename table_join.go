package framed

type TableJoinType int

const (
	TableJoinInner TableJoinType = iota
	TableJoinLeft
	TableJoinRight
	TableJoinCross
)

type TableJoinConfig struct {
	JoinType TableJoinType
	Left     string
	Right    string
	Rename   string
}

func (c *TableJoinConfig) From(s string) *TableJoinConfig {
	c.Left = s
	return c
}

func (c *TableJoinConfig) To(s string) *TableJoinConfig {
	c.Right = s
	return c
}

func (c *TableJoinConfig) As(s string) *TableJoinConfig {
	c.Rename = s
	return c
}

func JoinConfig(tp TableJoinType, col string) *TableJoinConfig {
	return &TableJoinConfig{
		JoinType: tp,
		Left:     col,
		Right:    col,
	}
}

func Join(tb1 *Table, tb2 *Table, config *TableJoinConfig) (*Table, error) {
	allCols := make([]string, 0)
	allCols = append(allCols, tb1.State.Columns...)
	allCols = append(allCols, tb2.State.Columns...)

	columns := SliceUnique(allCols)
	cLen := len(columns)
	lIndex := tb1.State.Index(config.Left)
	rIndex := tb2.State.Index(config.Right)

	if lIndex < 0 {
		return nil, UnknownColumnError(config.Left)
	} else if rIndex < 0 {
		return nil, UnknownColumnError(config.Right)
	}

	lIndexes := tb1.State.ResolveIndexes(tb1.State.Columns)
	rIndexes := tb2.State.ResolveIndexes(SliceOmit(tb2.State.Columns, []string{config.Right}))

	df := New().MarkResolved()
	df.UseColumns(columns)

	for _, col := range columns {
		if tb1.State.HasDefinition(col) {
			df.SetDefinition(col, tb1.State.Definition(col))
		} else {
			df.SetDefinition(col, tb2.State.Definition(col))
		}
	}

	idx := 0
	switch config.JoinType {
	case TableJoinInner:
		for _, row := range tb1.Rows {
			jVal := row.At(lIndex)
			mRow := tb2.Find(func(r *Row) bool {
				return jVal == r.At(rIndex)
			})

			if mRow != nil {
				data := make([]any, 0, cLen)
				data = append(data, row.Pick(lIndexes)...)
				data = append(data, mRow.Pick(rIndexes)...)

				df.AddRow(&Row{
					Index:   idx,
					Columns: data,
				})
				idx++
			}
		}
	case TableJoinLeft:
		for _, row := range tb1.Rows {
			jVal := row.At(lIndex)
			mRow := tb2.Find(func(r *Row) bool {
				return jVal == r.At(rIndex)
			})

			nRow := NewRow(idx, cLen)
			data := make([]any, 0, cLen)
			data = append(data, row.Pick(lIndexes)...)
			if mRow != nil {
				data = append(data, mRow.Pick(rIndexes)...)
			}

			df.AddRow(nRow)
			idx++
		}
	}

	return df, nil
}
