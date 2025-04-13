package framed

// pipeline action to rename table column(s)
type ActionRenameColumn struct {
	pairs [][]string
}

func (a ActionRenameColumn) ExecName() string {
	return "rename_column"
}

func (a ActionRenameColumn) Execute(src *Table) (*Table, error) {
	df := src.Clone()

	for _, pair := range a.pairs {
		idx := df.State.Index(pair[0])
		if idx > -1 {
			df.State.Columns[idx] = pair[1]
			df.State.Indexes[pair[1]] = idx
			df.State.Definitions[pair[1]] = df.State.Definition(pair[0])

			delete(df.State.Indexes, pair[0])
			delete(df.State.Definitions, pair[0])
		}
	}

	return df, nil
}

func RenameColumn(name string, newName string) *ActionRenameColumn {
	return &ActionRenameColumn{
		pairs: [][]string{
			{name, newName},
		},
	}
}

func RenameColumns(pairs ...[]string) *ActionRenameColumn {
	return &ActionRenameColumn{
		pairs: pairs,
	}
}
