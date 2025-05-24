package framed_test

import (
	"fmt"
	"testing"

	"github.com/OpenRunic/framed"
)

func TestActionAddColumn(t *testing.T) {
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.AddColumn("name", "", func(s *framed.State, r *framed.Row) (string, error) {
			return fmt.Sprintf("%s %s", r.At(s.Index("last_name")), r.At(s.Index("first_name"))), nil
		}),
	)

	if err != nil {
		t.Error(err)
	} else {
		if !newDF.State.HasColumn("name") {
			t.Error("expected new column in table but found none")
		}
	}
}
