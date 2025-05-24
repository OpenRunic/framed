package framed_test

import (
	"testing"

	"github.com/OpenRunic/framed"
)

func TestActionChangeColumnType(t *testing.T) {
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.ChangeColumnType("age", "", func(_ *framed.State, _ *framed.Row, a any) (string, error) {
			v := a.(int64)
			if v > 18 {
				return "adult", nil
			} else if v > 12 {
				return "teen", nil
			}

			return "kid", nil
		}),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.State.DataType("age") != framed.ToType("") {
			t.Error("expected column in table to be of type string")
		}
	}
}
