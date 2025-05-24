package framed

import (
	"reflect"
)

type ChangeColumnDataReader[T any] = func(*State, *Row, any) (T, error)

type ActionChangeColumnType[T any] struct {
	column   string
	dataType reflect.Type
	callback ChangeColumnDataReader[T]
}

func (a ActionChangeColumnType[T]) ActionName() string {
	return "change_column_type"
}

func (a ActionChangeColumnType[T]) Execute(src *Table) (*Table, error) {
	df := src.CloneE()
	colIndex := df.State.Index(a.column)

	typeChanged := a.dataType != df.State.DataType(a.column)
	if typeChanged {
		df.SetDefinition(a.column, NewDefinition(a.dataType))
	}

	for _, r := range src.Rows {
		row := r.Clone()
		val, err := a.callback(df.State, row, row.At(colIndex))
		if err != nil {
			return nil, err
		}

		df.AddRow(row.Set(colIndex, val))
	}

	return df, nil
}

// ChangeColumnType updates columns on every rows and resolves
// column value using callback and generates new table.
// Generic type of [T any] is applied.
//
//	$0 : Column Name (string)
//	$1 : Sample Value (T)
//	$2 : func(*framed.State, *framed.Row, any) (T, error)
//
//	newTable, err := table.Execute(
//		...
//		framed.ChangeColumnType($0, $1, $2),
//		...
//	)
func ChangeColumnType[T any](column string, sample T, cb ChangeColumnDataReader[T]) *ActionChangeColumnType[T] {
	return &ActionChangeColumnType[T]{
		column:   column,
		dataType: ToType(sample),
		callback: cb,
	}
}
