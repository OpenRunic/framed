package framed_test

import (
	"reflect"
	"testing"
)

func TestTableRowPutValue(t *testing.T) {
	colIndex := 3
	df := SampleTestTable(t)
	row := df.First()
	def := df.State.DefinitionAt(colIndex)
	err := row.Patch(def, colIndex, "")

	if err == nil {
		t.Error("expected error on invalid value set, got none")
	} else {
		r, err := row.Pick(df.State, "age")
		if err != nil {
			t.Error(err)
		} else {
			if reflect.TypeOf(r.At(0)).Kind() != reflect.Int32 {
				t.Error("failed to clone and pick column from existing row")
			}
		}
	}
}
