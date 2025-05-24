package framed

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
)

// CountNull allows you to count empty/nil values
func (t *Table) CountNull() map[string]int {
	counts := make(map[string]int)
	for _, col := range t.State.Columns {
		counts[col] = 0
	}

	for _, row := range t.Rows {
		for idx, col := range t.State.Columns {
			if IsEmpty(t.State.Definitions[col], row.At(idx)) {
				counts[col]++
			}
		}
	}

	return counts
}

// FillNull allows you to replace empty/nil values with defined value
func (t *Table) FillNull(values map[string]any) error {
	for column, value := range values {
		index := t.State.Index(column)
		if index < 0 {
			return UnknownColumnError(column)
		}

		def := t.State.Definition(column)
		if def.Type != reflect.TypeOf(value) {
			return InvalidColumnValueTypeError(column, def.Type)
		}

		for _, row := range t.Rows {
			colVal := row.At(index)
			if IsEmpty(def, colVal) {
				row.Set(index, value)
			}
		}
	}

	return nil
}

// RunNumeric runs numeric operation on column values
func (t *Table) RunNumeric(column string, cb func(*Row, float64)) error {
	if t.IsEmpty() {
		return nil
	}

	index := t.State.Index(column)
	def := t.State.Definition(column)
	if def == nil {
		return UnknownColumnError(column)
	}

	resolved := ""
	if !slices.Contains(NumericKinds, def.Kind()) {
		v := t.First().At(index)
		_, ok := v.(NumericReader)
		if ok {
			resolved = "iface"
		} else {
			resolved = "err"
		}
	}

	if resolved == "err" {
		return NonNumericColumnValueError(column)
	}

	for _, row := range t.Rows {
		var val float64
		if resolved == "" {
			v, _ := strconv.ParseFloat(fmt.Sprint(row.At(index)), 64)
			val = v
		} else {
			v, _ := row.At(index).(NumericReader)
			val += v.NumericRead()
		}

		cb(row, val)
	}

	return nil
}

// Avg calculates average value of column
func (t *Table) Avg(column string) (float64, error) {
	var total float64 = 0

	err := t.RunNumeric(column, func(r *Row, f float64) {
		total += f
	})
	if err != nil {
		return 0, err
	}

	return total / float64(t.Length()), nil
}

// Sum calculates total sum of column
func (t *Table) Sum(column string) (float64, error) {
	var total float64 = 0

	err := t.RunNumeric(column, func(r *Row, f float64) {
		total += f
	})
	if err != nil {
		return 0, err
	}

	return total, nil
}

// Col picks all the values from column
func (t *Table) Col(column string) []any {
	index := t.State.Index(column)
	values := make([]any, t.Length())

	if index > -1 {
		for idx, row := range t.Rows {
			values[idx] = row.At(index)
		}
	}

	return values
}

// GetUnique collects unique values by column
func (t *Table) GetUnique(column string) []any {
	index := t.State.Index(column)
	if index < 0 {
		return make([]any, 0)
	}

	m := map[any]bool{}
	return ReduceTable(t, []any{}, func(res []any, row *Row) []any {
		val := row.At(index)

		if !m[val] {
			m[val] = true
			res = append(res, val)
		}

		return res
	})
}

// CountValues counts unique value totals by column
func (t *Table) CountValues(column string) []ColumnStatValue {
	index := t.State.Index(column)
	if index < 0 {
		return make([]ColumnStatValue, 0)
	}

	m := map[any]bool{}
	res := ReduceTable(t, map[any]int{}, func(res map[any]int, row *Row) map[any]int {
		val := row.At(index)

		if !m[val] {
			m[val] = true
			res[val] = 0
		}

		res[val]++

		return res
	})

	v := make([]ColumnStatValue, 0)
	for key, count := range res {
		v = append(v, ColumnStatValue{
			Key:   key,
			Value: count,
		})
	}

	return v
}
