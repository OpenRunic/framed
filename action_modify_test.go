package framed_test

import (
	"testing"

	"github.com/DecxBase/framed"
)

func TestTableModifyRows(t *testing.T) {
	inc := 100
	df := SampleTestTable(t)
	newDF, err := df.Execute(
		framed.ModifyRow(func(s *framed.State, r *framed.Row) *framed.Row {
			r.Set(0, framed.ColumnValue(r, 0, 0)+inc)
			return r
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
