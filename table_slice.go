package framed

import (
	"iter"
	"slices"
)

// Slice retrieves slice of rows at given start and end
func (t *Table) Slice(start int, end int) []*Row {
	return t.Rows[start:end]
}

// RSlice retrieves slice of rows at given start and end but reverse
func (t *Table) RSlice(start int, end int) []*Row {
	rows := make([]*Row, t.Length())
	_ = copy(rows, t.Rows)
	slices.Reverse(rows)

	return rows[start:end]
}

// First returns the first [Row] of table
func (t *Table) First() *Row {
	return t.At(0)
}

// Last returns the last [Row] of table
func (t *Table) Last() *Row {
	return t.At(t.Length() - 1)
}

// Chunk build new table from provided slice indexes
func (t *Table) Chunk(start int, end int) *Table {
	df := t.CloneE()
	rows := t.Slice(start, end)

	for idx, r := range rows {
		df.AddRow(r.Clone().WithIndex(idx))
	}

	return df
}

// Chunks split table into multiple chunk of tables from provided slice indexes
func (t *Table) Chunks(limit int) iter.Seq2[int, *Table] {
	return func(yield func(int, *Table) bool) {
		count := t.Length()
		pages := count / limit
		if count > (pages * limit) {
			pages++
		}

		for i := range pages {
			start := i * limit
			end := min(start+limit, count)

			if !yield(i, t.Chunk(start, end)) {
				return
			}
		}
	}
}
