package framed_test

import (
	"testing"

	"github.com/DecxBase/framed"
)

func TestTableChangeType(t *testing.T) {
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.ChangeType("age", "", func(_ *framed.State, _ *framed.Row, a any) string {
			v := a.(int32)
			if v > 18 {
				return "adult"
			} else if v > 12 {
				return "teen"
			}

			return "kid"
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
