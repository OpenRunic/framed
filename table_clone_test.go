package framed_test

import (
	"reflect"
	"testing"
)

func TestTableClone(t *testing.T) {
	df := SampleTestTable()
	cloned := df.Clone()

	if !reflect.DeepEqual(df, cloned) {
		t.Error("unable to produce identical clone of table")
	}
}

func TestTableCloneOptionsDefinitions(t *testing.T) {
	df := SampleTestTable()
	cloned := df.CloneE().MarkUnresolved()

	if !reflect.DeepEqual(df.State, cloned.State) {
		t.Error("unable to produce identical clone state of table")
	}
	if !reflect.DeepEqual(df.Options, cloned.Options) {
		t.Error("unable to produce identical clone options of table")
	}
	if cloned.IsResolved() || reflect.DeepEqual(df.Rows, cloned.Rows) {
		t.Error("unable to detect empty rows in table")
	}
}

func TestTableCloneWithoutRows(t *testing.T) {
	df := SampleTestTable()
	cloned := df.CloneE()

	if !reflect.DeepEqual(df.State, cloned.State) {
		t.Error("unable to produce identical clone state of table")
	}
	if !reflect.DeepEqual(df.Options, cloned.Options) {
		t.Error("unable to produce identical clone options of table")
	}
	if !cloned.IsResolved() || reflect.DeepEqual(df.Rows, cloned.Rows) {
		t.Error("unable to detect empty rows in table")
	}
}
