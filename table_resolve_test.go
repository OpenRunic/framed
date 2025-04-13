package framed_test

import (
	"reflect"
	"testing"
)

func TestTableResolveDefinition(t *testing.T) {
	df := SampleTestTable()
	def := df.ResolveValueDefinition(20, "test", "2.501")

	if def.Kind() != reflect.Float64 {
		t.Error("failed to resolve valid data type")
	}
}
