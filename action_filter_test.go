package framed_test

import (
	"testing"

	"github.com/OpenRunic/framed"
)

func TestTableFilterRows(t *testing.T) {
	limit := 10
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.FilterRow(func(_ *framed.State, r *framed.Row) bool {
			return r.Index < limit
		}),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.Length() > limit {
			t.Error("expected rows to be filtered as provided")
		}
	}
}
