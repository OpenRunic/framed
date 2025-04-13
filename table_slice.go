package framed

import (
	"iter"
	"slices"
)

// access slice of rows
func (t *Table) Slice(start int, end int) []*Row {
	return t.Rows[start:end]
}

// access slice of rows in reverse
func (t *Table) RSlice(start int, end int) []*Row {
	rows := make([]*Row, t.Length())
	_ = copy(rows, t.Rows)
	slices.Reverse(rows)

	return rows[start:end]
}

func (t *Table) First() *Row {
	return t.At(0)
}

func (t *Table) Last() *Row {
	return t.At(t.Length() - 1)
}

func (t *Table) Head(limit int) []*Row {
	return t.Slice(0, limit)
}

func (t *Table) Tail(limit int) []*Row {
	return t.RSlice(0, limit)
}

// build new table from provided slice indexes
func (t *Table) Chunk(start int, end int) *Table {
	df := t.CloneE()
	rows := t.Slice(start, end)

	for idx, r := range rows {
		df.AddRow(r.Clone().WithIndex(idx))
	}

	return df
}

// split table into multiple chunk of tables from provided slice indexes
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
