package framed

import (
	"io"
	"iter"
)

// use reader to load data into table
func (t *Table) Read(r io.Reader) error {
	var n int
	var err error
	buf := make([]byte, 256)

	for err == nil {
		n, err = r.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
		} else {
			if t.IsAtMaxLine() {
				break
			}

			err := t.Insert(buf[:n])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// convert slice of values to column slices
func (t *Table) SliceToColumns(values []string) ([]any, error) {
	hLen := t.ColLength()
	columns := make([]any, len(values))

	for idx, value := range values {
		if hLen > idx {
			c, err := t.AsColumn(idx, value)
			if err != nil {
				return nil, err
			}
			columns[idx] = c
		}
	}

	return columns, nil
}

// make column info for provided value
func (t *Table) AsColumn(idx int, value string) (any, error) {
	def := t.State.Definition(t.State.ColumnName(idx))
	val, err := ColumnValueDecoder(def, value)
	if err != nil {
		return nil, ColError(t.Length()-1, idx, t.State.ColumnName(idx), err, "value_decode")
	}
	return val, nil
}

func (t *Table) AddRow(rows ...*Row) {
	t.Rows = append(t.Rows, rows...)
}

// add new column info to state
func (t *Table) AppendColumn(pos int, name string) {
	t.State.Indexes[name] = pos
	t.State.Columns = append(t.State.Columns, name)
}

// insert byte data as row
func (t *Table) Insert(b []byte) error {
	return t.InsertLine(string(b))
}

// insert string line as row
func (t *Table) InsertLine(line string) error {
	return t.InsertSlice(SplitAtChar(line, byte(t.Options.Separator)))
}

// insert slice of strings as row
func (t *Table) InsertSlice(values []string) error {
	line := t.Length()
	if line == 0 {
		if t.Options.IgnoreHeader {
			return nil
		}

		if t.State.IsEmpty() {
			t.UseColumns(values)
			return nil
		}
	}

	if !t.resolved {
		err := t.ResolveTypes(t.State.Columns, values)
		if err != nil {
			return err
		}
	}

	columns, err := t.SliceToColumns(values)
	if err != nil {
		return err
	}

	t.AddRow(&Row{
		Index:   line,
		Columns: columns,
	})

	return nil
}

// insert list of lines
func (t *Table) InsertLines(lines []string) error {
	for _, line := range lines {
		if t.IsAtMaxLine() {
			break
		}

		err := t.InsertLine(line)
		if err != nil {
			return err
		}
	}

	return nil
}

// insert list of slices
func (t *Table) InsertSlices(ss [][]string) error {
	for _, s := range ss {
		if t.IsAtMaxLine() {
			break
		}

		err := t.InsertSlice(s)
		if err != nil {
			return err
		}
	}

	return nil
}

// insert list of bytes from provided iterator
func (t *Table) InsertGenBytes(it iter.Seq[[]byte]) error {
	for b := range it {
		if t.IsAtMaxLine() {
			break
		}

		err := t.InsertLine(string(b))
		if err != nil {
			return err
		}
	}

	return nil
}
