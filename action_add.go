package framed

import (
	"reflect"
)

type AddColumnDataReader[T any] = func(*State, *Row) T

// pipeline action to add new column to table
type ActionAddColumn[T any] struct {
	name     string
	dataType reflect.Type
	callback AddColumnDataReader[T]
}

func (a ActionAddColumn[T]) ExecName() string {
	return "add_column"
}

func (a ActionAddColumn[T]) Execute(src *Table) (*Table, error) {
	df := src.CloneE()
	pos := src.ColLength()

	df.ResolveDefinition(pos, a.name, a.dataType)
	df.AppendColumn(pos, a.name)

	idx := 0
	for _, r := range src.Rows {
		row := r.Clone().WithIndex(idx)
		df.AddRow(row.AddColumn(a.name, a.callback(df.State, row)))
		idx += 1
	}

	return df, nil
}

func AddColumn[T any](name string, sample T, cb AddColumnDataReader[T]) *ActionAddColumn[T] {
	return &ActionAddColumn[T]{
		name:     name,
		dataType: ToType(sample),
		callback: cb,
	}
}
