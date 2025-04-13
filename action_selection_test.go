package framed_test

import (
	"reflect"
	"testing"

	"github.com/DecxBase/framed"
)

func TestTablePickColumns(t *testing.T) {
	cols := []string{"first_name", "age"}
	df := SampleTestTable(t)
	newDF, err := df.Execute(
		framed.PickColumn(cols...),
	)

	if err != nil {
		t.Error(err)
	} else {
		if !reflect.DeepEqual(newDF.State.Columns, cols) {
			t.Error("failed to pick columns")
		}
	}
}

func TestTableDropColumns(t *testing.T) {
	df := SampleTestTable(t)
	newDF, err := df.Execute(
		framed.DropColumn("id", "last_name"),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.State.HasColumn("id") || newDF.State.HasColumn("last_name") {
			t.Error("failed to drop columns")
		}
	}
}
