package framed

import (
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

// At gives access to column at n index
func (r *Row) At(n int) any {
	if n > -1 {
		return r.Columns[n]
	}
	return nil
}

// Set updates column value at n index
func (r *Row) Set(n int, value any) *Row {
	r.Columns[n] = value
	return r
}

// Patch attempts to update column value at n index
// and throws error on type fail
func (r *Row) Patch(def *Definition, n int, value any) error {
	tp := ToType(value)
	if def.Type != tp {
		return WriteColumnValueError(r.Index, n, def.Type, tp)
	}

	r.Set(n, value)
	return nil
}

// Encode attempts to encode column value to string
func (r *Row) Encode(def *Definition, n int) (string, error) {
	colVal := r.At(n)
	if IsEmpty(def, colVal) {
		return "", nil
	}

	val, err := ColumnValueEncoder(def, colVal)
	if err != nil {
		return "", EncodeColumnValueError(r.Index, n, err)
	}
	return val, nil
}

// AsSlice encodes columns to slice of strings or throws error
func (r *Row) AsSlice(s *State) ([]string, error) {
	colCount := len(r.Columns)
	values := make([]string, colCount)

	for i := range colCount {
		val, err := r.Encode(s.DefinitionAt(i), i)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}

	return values, nil
}

// Copy copies from value slice to their index
func (r *Row) Copy(v []any) *Row {
	cap := len(r.Columns)
	for i, v := range v {
		if cap > i {
			r.Columns[i] = v
		}
	}
	return r
}

// Pick selects provided columns from the row or throws error
func (r *Row) Pick(indexes []int) []any {
	columns := make([]any, 0)

	for _, idx := range indexes {
		if idx > -1 {
			columns = append(columns, r.At(idx))
		} else {
			columns = append(columns, nil)
		}
	}

	return columns
}

// Clone duplicates the row as-is
func (r *Row) Clone() *Row {
	return &Row{
		Index:   r.Index,
		Columns: slices.Clone(r.Columns),
	}
}

// CloneP create duplicate row from selected columns
func (r *Row) CloneP(indexes []int) *Row {
	return &Row{
		Index:   r.Index,
		Columns: r.Pick(indexes),
	}
}

// NewRow creates a row with defined length
func NewRow(idx int, size int) *Row {
	return &Row{
		Index:   idx,
		Columns: make([]any, size),
	}
}
