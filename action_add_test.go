package framed_test

import (
	"fmt"
	"testing"

	"github.com/DecxBase/framed"
)

func TestTableAddColumn(t *testing.T) {
	df := SampleTestTable(t)
	newDF, err := df.Execute(
		framed.AddColumn("name", "", func(s *framed.State, r *framed.Row) string {
			return fmt.Sprintf("%s %s", r.At(s.Index("last_name")), r.At(s.Index("first_name")))
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
