package framed_test

import (
	"testing"

	"github.com/DecxBase/framed"
)

func TestTableRenameColumn(t *testing.T) {
	df := SampleTestTable(t)
	newDF, err := df.Execute(
		framed.RenameColumn("age", "just_a_num"),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.State.HasColumn("age") {
			t.Error("failed to rename table column")
		}
	}
}
