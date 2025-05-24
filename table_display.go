package framed

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

// String generates printable table data
func (t *Table) String() string {
	return t.PreparePrint(t.Rows, true)
}

// Head returns the first n [Row] of table
func (t *Table) Head(limit int) string {
	return t.PreparePrint(t.Slice(0, limit), false)
}

// Tail returns the last n [Row] of table
func (t *Table) Tail(limit int) string {
	return t.PreparePrint(t.RSlice(0, limit), false)
}

// PreparePrint takes slice of [Row] and displays pretty info about them
func (t *Table) PreparePrint(rows []*Row, useLimit bool) string {
	cLen := t.ColLength()
	if cLen < 1 || !t.resolved {
		return t.Options.EmptyText
	}

	rLen := len(rows)
	cols := t.State.Columns
	colTypes := make(map[int]reflect.Type, cLen)
	colNames := make(map[int]string, cLen)
	sampleSize := min(t.Options.SampleSize, rLen)
	sampleMaxLength := t.Options.SampleMaxLength

	if !useLimit {
		sampleSize = rLen
	}

	var b strings.Builder
	colWidths := make(map[int]int)
	values := make(map[int][]string, len(cols))

	// go through each columns
	for idx, col := range cols {
		def := t.State.Definition(col)
		tLen := len(def.Type.String())

		colNames[idx] = col
		colTypes[idx] = def.Type
		colWidths[idx] = max(len(col), tLen)

		// collect column samples
		samples := make([]string, sampleSize)
		for si := range sampleSize {
			val, err := rows[si].Encode(def, idx)

			if err != nil {
				samples[si] = t.Options.EmptyText
			} else {
				if len(val) > 0 {
					if useLimit && tLen < len(val) && len(val) > sampleMaxLength {
						samples[si] = val[0:(sampleMaxLength-3)] + "..."
					} else {
						samples[si] = val
					}
				} else {
					samples[si] = t.Options.EmptyText
				}

				if len(samples[si]) > colWidths[idx] {
					colWidths[idx] = len(samples[si])
				}
			}
		}

		values[idx] = samples
	}

	// print shape
	b.WriteString(fmt.Sprintf("\nshape: %d, %d\n", rLen, cLen))

	// check has table name
	if len(t.Options.Name) > 0 {
		fullWidth := MapReduce(colWidths, 0, func(w int, i int, cW int) int {
			return w + cW + 2
		}) + cLen - 1

		// title print start
		b.WriteRune(t.Options.PrettyFormatCharacters[0])
		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[12]), fullWidth))
		b.WriteRune(t.Options.PrettyFormatCharacters[2])
		b.WriteString("\n")

		// centered table title
		space := "\u0020"
		n := (fullWidth - utf8.RuneCountInString(t.Options.Name)) / 2
		b.WriteString(fmt.Sprintf(
			"%s%s%s%s%s\n",
			string(t.Options.PrettyFormatCharacters[3]),
			strings.Repeat(space, n),
			t.Options.Name,
			strings.Repeat(space, n+1),
			string(t.Options.PrettyFormatCharacters[3]),
		))

		// title print end
		b.WriteRune(t.Options.PrettyFormatCharacters[9])
		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[12]), fullWidth))
		b.WriteRune(t.Options.PrettyFormatCharacters[11])
		b.WriteString("\n")
	}

	// add table start
	b.WriteRune(t.Options.PrettyFormatCharacters[0])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[1])
		}

		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[12]), colWidths[i]+2))
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[2])
	b.WriteString("\n")

	// add column names
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[4])
		}

		b.WriteRune(t.Options.PrettyFormatCharacters[14])
		b.WriteString(fmt.Sprintf("%-*s", colWidths[i], colNames[i]))
		b.WriteRune(t.Options.PrettyFormatCharacters[14])
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	b.WriteString("\n")

	// add column names separator
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[4])
		}

		b.WriteRune(t.Options.PrettyFormatCharacters[14])
		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[13]), colWidths[i]))
		b.WriteRune(t.Options.PrettyFormatCharacters[14])
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	b.WriteString("\n")

	// add column types
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[4])
		}

		b.WriteRune(t.Options.PrettyFormatCharacters[14])
		b.WriteString(fmt.Sprintf("%-*s", colWidths[i], colTypes[i]))
		b.WriteRune(t.Options.PrettyFormatCharacters[14])
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[3])
	b.WriteString("\n")

	// add body start
	b.WriteRune(t.Options.PrettyFormatCharacters[5])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[6])
		}
		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[7]), colWidths[i]+2))
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[8])
	b.WriteString("\n")

	// add sample rows
	for si := range sampleSize {
		b.WriteRune(t.Options.PrettyFormatCharacters[3])
		for ci := range cLen {
			if ci > 0 {
				b.WriteRune(t.Options.PrettyFormatCharacters[4])
			}

			b.WriteRune(t.Options.PrettyFormatCharacters[14])
			b.WriteString(fmt.Sprintf("%-*s", colWidths[ci], values[ci][si]))
			b.WriteRune(t.Options.PrettyFormatCharacters[14])
		}
		b.WriteRune(t.Options.PrettyFormatCharacters[3])
		b.WriteString("\n")
	}

	// add table end
	b.WriteRune(t.Options.PrettyFormatCharacters[9])
	for i := range cLen {
		if i > 0 {
			b.WriteRune(t.Options.PrettyFormatCharacters[10])
		}

		b.WriteString(strings.Repeat(string(t.Options.PrettyFormatCharacters[12]), colWidths[i]+2))
	}
	b.WriteRune(t.Options.PrettyFormatCharacters[11])

	return b.String()
}
