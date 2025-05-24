package framed

import (
	"reflect"
)

type AddColumnDataReader[T any] = func(*State, *Row) (T, error)

type ActionAddColumn[T any] struct {
	column   string
	dataType reflect.Type
	callback AddColumnDataReader[T]
}

func (a ActionAddColumn[T]) ActionName() string {
	return "add_column"
}

func (a ActionAddColumn[T]) Execute(src *Table) (*Table, error) {
	df := src.CloneE()
	pos := src.ColLength()

	df.ResolveDefinition(a.column, a.dataType)
	df.AppendColumn(pos, a.column)

	for _, r := range src.Rows {
		row := r.Clone()
		v, err := a.callback(df.State, row)
		if err != nil {
			return nil, err
		}

		df.AddRow(row.AddColumn(v))
	}

	return df, nil
}

// AddColumn adds new column to every rows and resolves
// column value using callback and generates new table.
// Generic type of [T any] is applied.
//
//	$0 : Column Name (string)
//	$1 : Sample Value (T)
//	$2 : func(*framed.State, *framed.Row) (T, error)
//
//	newTable, err := table.Execute(
//		...
//		framed.AddColumn($0, $1, $2),
//		...
//	)
func AddColumn[T any](column string, sample T, cb AddColumnDataReader[T]) *ActionAddColumn[T] {
	return &ActionAddColumn[T]{
		column:   column,
		dataType: ToType(sample),
		callback: cb,
	}
}
