package framed_test

import (
	"bytes"
	"io"
	"testing"
)

func TestTableToWriter(t *testing.T) {
	df := SampleTestTable(t)

	var b bytes.Buffer
	bWriter := io.Writer(&b)

	err := df.Write(bWriter)
	if err != nil {
		t.Error(err)
	} else if b.Len() == 0 {
		t.Error("failed to write table to writer")
	}
}
