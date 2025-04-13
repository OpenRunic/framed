package framed

import (
	"reflect"
)

type ChangeColumnDataReader[T any] = func(*State, *Row, any) T

// pipeline action to change column type on table
type ActionChangeColumnType[T any] struct {
	name     string
	dataType reflect.Type
	callback ChangeColumnDataReader[T]
}

func (a ActionChangeColumnType[T]) ExecName() string {
	return "change_column_type"
}

func (a ActionChangeColumnType[T]) Execute(src *Table) (*Table, error) {
	df := src.CloneE()
	colIndex := df.State.Index(a.name)

	typeChanged := a.dataType != df.State.DataType(a.name)
	if typeChanged {
		df.SetDefinition(a.name, NewDefinition(a.dataType))
	}

	for _, r := range src.Rows {
		row := r.Clone()
		df.AddRow(row.Set(colIndex, a.callback(df.State, row, row.At(colIndex))))
	}

	return df, nil
}

func ChangeType[T any](name string, sample T, cb ChangeColumnDataReader[T]) *ActionChangeColumnType[T] {
	return &ActionChangeColumnType[T]{
		name:     name,
		dataType: ToType(sample),
		callback: cb,
	}
}
