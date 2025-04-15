package framed

import (
	"io"
	"os"
)

// Write writes table data to [io.Writer]
func (t *Table) Write(f io.Writer) error {
	_, err := f.Write([]byte(JoinAtChar(t.State.Columns, t.Options.Separator)))
	if err != nil {
		return err
	}

	for _, r := range t.Rows {
		line, err := r.AsSlice(t.State)
		if err != nil {
			return err
		}

		_, err = f.Write([]byte("\n" + JoinAtChar(line, t.Options.Separator)))
		if err != nil {
			return err
		}
	}

	return nil
}

// Save saves the table data to provided file path
func (t *Table) Save(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Write(f)
	if err != nil {
		return err
	}

	return nil
}
