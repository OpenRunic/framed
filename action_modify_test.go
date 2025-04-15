package framed_test

import (
	"testing"

	"github.com/OpenRunic/framed"
)

func TestTableModifyRows(t *testing.T) {
	inc := 100
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.ModifyRow(func(_ *framed.State, r *framed.Row) *framed.Row {
			return r.Set(0, framed.ColumnValue(r, 0, 0)+inc)
		}),
	)

	if err != nil {
		t.Error(err)
	} else {
		fid := framed.ColumnValue(newDF.First(), 0, 0)
		if fid < inc {
			t.Error("expected rows to be modified as provided")
		}
	}
}
