package framed

import (
	"fmt"
	"slices"
)

// table row data
type Row struct {

	// index of row in table
	Index int

	// slice of column data
	Columns []any
}

func (r *Row) WithIndex(idx int) *Row {
	r.Index = idx
	return r
}

func (r *Row) AddColumn(value any) *Row {
	r.Columns = append(r.Columns, value)
	return r
}

func (r *Row) At(idx int) any {
	return r.Columns[idx]
}

func (r *Row) Set(idx int, value any) *Row {
	r.Columns[idx] = value
	return r
}

func (r *Row) Patch(def *ColumnDefinition, idx int, value any) error {
	tp := ToType(value)
	if def.Type != tp {
		return ColError(
			r.Index, idx, "",
			fmt.Errorf("set column value failed; %s != %s", def.Type, tp),
			"write_value",
		)
	}

	r.Set(idx, value)
	return nil
}

// encode the slice of any as slice of strings
func (r *Row) AsSlice(s *State) ([]string, error) {
	colCount := len(r.Columns)
	values := make([]string, colCount)

	for i := range colCount {
		val, err := ColumnValueEncoder(s.DefinitionAt(i), r.At(i))
		if err != nil {
			return nil, ColError(r.Index, i, s.ColumnName(i), err, "value_encode")
		}
		values[i] = val
	}

	return values, nil
}

// pick only provided columns from the row
func (r *Row) Pick(s *State, names ...string) ([]any, error) {
	cols := slices.Clone(names)
	if len(cols) < 1 {
		cols = slices.Clone(s.Columns)
	}

	columns := make([]any, 0)

	for _, col := range cols {
		idx := s.Index(col)
		if idx > -1 {
			columns = append(columns, r.At(idx))
		} else {
			return nil, RowError(r.Index, fmt.Errorf("unknown column [%s]", col), "not_found")
		}
	}

	return columns, nil
}

// clone the row with only selected columns
func (r *Row) CloneP(s *State, names ...string) (*Row, error) {
	columns, err := r.Pick(s, names...)
	if err != nil {
		return nil, err
	}

	return &Row{
		Index:   r.Index,
		Columns: columns,
	}, nil
}

// clone the row as-is
func (r *Row) Clone() *Row {
	return &Row{
		Index:   r.Index,
		Columns: slices.Clone(r.Columns),
	}
}
