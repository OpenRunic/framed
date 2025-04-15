package framed

import "fmt"

type ActionRenameColumn struct {
	pairs [][]string
}

func (a ActionRenameColumn) ActionName() string {
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
		} else {
			return nil, fmt.Errorf("unable to locate column: %s", pair[0])
		}
	}

	return df, nil
}

// RenameColumn renames a column generates the new table.
//
//	$0 : Existing Column Name
//	$1 : New Column Name
//
//	newTable, err := table.Execute(
//		...
//		framed.RenameColumn($0, $1),
//		...
//	)
func RenameColumn(name string, newName string) *ActionRenameColumn {
	return &ActionRenameColumn{
		pairs: [][]string{
			{name, newName},
		},
	}
}

// RenameColumns renames multiple columns at once and
// generates the new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.RenameColumns([]string{"EXISTING_COL_NAME", "NEW_COLUMN_NAME"}, ...),
//		...
//	)
func RenameColumns(pairs ...[]string) *ActionRenameColumn {
	return &ActionRenameColumn{
		pairs: pairs,
	}
}
