package framed_test

import (
	"fmt"
	"testing"

	"github.com/OpenRunic/framed"
)

func TestTableSelect(t *testing.T) {
	df := SampleTestTable()
	newDF, err := df.Execute(
		framed.Select(
			framed.ColAll("last_name", "age").Prefix("T_"),
			framed.Col("age").Alias("aged"),
			framed.Col("nidx").Modify(func(s *framed.State, r *framed.Row) (any, error) {
				return len(fmt.Sprint(r.At(1), r.At(2))), nil
			}),
			framed.Col("ranking").Build(func(_ *framed.Table) ([]any, error) {
				return []any{
					"top",
				}, nil
			}),
		),
	)

	if err != nil {
		t.Error(err)
	} else {
		if newDF.ColLength() != 7 {
			t.Errorf("failed to resolve new columns size")
		} else {
			idx := newDF.State.Index("ranking")
			if framed.ColumnValue(newDF.First(), idx, "") != "top" {
				t.Errorf("failed to assign value for first row's new column")
			}

			if framed.ColumnValue(newDF.At(1), idx, "") != "" {
				t.Errorf("failed to detect empty value for second row's new column")
			}
		}
	}
}
