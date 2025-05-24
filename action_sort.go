package framed

import (
	"sort"
)

type ActionSort struct {
	column         string
	callback       func(*State, *Row, *Row) bool
	columnCallback func(*State, any, any) bool
}

func (a ActionSort) ActionName() string {
	return "sort"
}

func (a ActionSort) Execute(src *Table) (*Table, error) {
	if src.IsEmpty() {
		return src.CloneE(), nil
	}

	df := src.Clone()

	if len(a.column) > 0 {
		index := df.State.Index(a.column)
		if index < 0 {
			return nil, UnknownColumnError(a.column)
		}

		sort.Slice(df.Rows, func(i, j int) bool {
			return a.columnCallback(df.State, df.Rows[i].At(index), df.Rows[j].At(index))
		})
	} else {
		sort.Slice(df.Rows, func(i, j int) bool {
			return a.callback(df.State, df.Rows[i], df.Rows[j])
		})
	}

	return df, nil
}

// Sort iterates through rows and sorts the rows
// to build a new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.Sort(func(*framed.State, *framed.Row, *framed.Row) bool),
//		...
//	)
func Sort(cb func(*State, *Row, *Row) bool) *ActionSort {
	return &ActionSort{
		callback: cb,
	}
}

// SortBy iterates through rows and sorts the rows by column
// to build a new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.SortBy("col1", func(*framed.State, any, any) bool),
//		...
//	)
func SortBy(column string, cb func(*State, any, any) bool) *ActionSort {
	return &ActionSort{
		column:         column,
		columnCallback: cb,
	}
}
