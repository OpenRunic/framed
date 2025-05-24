package framed_test

import (
	"testing"

	"github.com/OpenRunic/framed"
)

func TestTableSort(t *testing.T) {
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.SortBy("age", func(s *framed.State, age1, age2 any) bool {
			return age2.(int64) > age1.(int64)
		}),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.First().At(3).(int64) > newDF.At(1).At(3).(int64) {
			t.Error("failed to sort table column")
		}
	}
}
