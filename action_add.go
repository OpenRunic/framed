package framed

import (
	"reflect"
)

type AddColumnDataReader[T any] = func(*State, *Row) T

type ActionAddColumn[T any] struct {
	name     string
	dataType reflect.Type
	callback AddColumnDataReader[T]
}

func (a ActionAddColumn[T]) ActionName() string {
	return "add_column"
}

func (a ActionAddColumn[T]) Execute(src *Table) (*Table, error) {
	df := src.CloneE()
	pos := src.ColLength()

	df.ResolveDefinition(a.name, a.dataType)
	df.AppendColumn(pos, a.name)

	for _, r := range src.Rows {
		row := r.Clone()
		df.AddRow(row.AddColumn(a.callback(df.State, row)))
	}

	return df, nil
}

// AddColumn adds new column to every rows and resolves
// column value using callback and generates new table.
// Generic type of [T any] is applied.
//
//	$0 : Column Name
//	$1 : Sample Value of T
//	$2 : func(*framed.State, *framed.Row) T
//
//	newTable, err := table.Execute(
//		...
//		framed.AddColumn($0, $1, $2),
//		...
//	)
func AddColumn[T any](name string, sample T, cb AddColumnDataReader[T]) *ActionAddColumn[T] {
	return &ActionAddColumn[T]{
		name:     name,
		dataType: ToType(sample),
		callback: cb,
	}
}
