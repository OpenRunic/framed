package framed

import (
	"fmt"
	"slices"
)

// Row holds columns of row in table
type Row struct {

	// index of row in table
	Index int

	// slice of column data
	Columns []any
}

// WithIndex updates the row index in table
func (r *Row) WithIndex(idx int) *Row {
	r.Index = idx
	return r
}

// AddColumn appends new column value to row
func (r *Row) AddColumn(value any) *Row {
	r.Columns = append(r.Columns, value)
	return r
}

// At gives access to column at x index
func (r *Row) At(idx int) any {
	return r.Columns[idx]
}

// Set updates column value at x index
func (r *Row) Set(idx int, value any) *Row {
	r.Columns[idx] = value
	return r
}

// Patch attempts to update column value at x index
// and throws error on type fail
func (r *Row) Patch(def *Definition, idx int, value any) error {
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

// AsSlice encodes columns to slice of strings or throws error
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

// Pick selects provided columns from the row or throws error
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

// CloneP create duplicate row from selected columns
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

// Clone duplicates the row as-is
func (r *Row) Clone() *Row {
	return &Row{
		Index:   r.Index,
		Columns: slices.Clone(r.Columns),
	}
}
