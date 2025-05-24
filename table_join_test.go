package framed_test

import (
	"fmt"
	"testing"

	"github.com/OpenRunic/framed"
)

func TestTableJoin(t *testing.T) {
	tb1 := SampleTestTable().Chunk(0, 5)

	tb2 := framed.New()
	tb2.SetDefinition("id", tb1.State.Definition("id"))
	tb2.SetDefinition("ssn", framed.NewDefinition(framed.ToType("")))
	tb2.UseColumns([]string{"id", "ssn"})
	tb2.MarkResolved()

	tb2.AddRows(framed.SliceMap(tb1.Col("id"), func(idx int, id any) *framed.Row {
		return &framed.Row{
			Index:   idx,
			Columns: []any{id.(int) + 1, fmt.Sprintf("%08v", id)},
		}
	}))

	joined, err := framed.Join(tb1, tb2, framed.JoinConfig(framed.TableJoinLeft, "id").To("id"))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(joined)
	}
}
