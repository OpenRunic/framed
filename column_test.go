package framed_test

import (
	"testing"

	"github.com/OpenRunic/framed"
)

var testColSplitValue = `aa,bb,cc,dd,"test value, data value"`

func TestColumnSplits(t *testing.T) {
	splits := framed.SplitAtChar(testColSplitValue, ',')

	if len(splits) != 5 {
		t.Error("failed to split columns as expected")
	}
}

func TestColumnValuesJoin(t *testing.T) {
	line := framed.JoinAtChar(
		[]string{"aa", "bb", "cc", "dd", "test value, data value"},
		',',
	)

	if line != testColSplitValue {
		t.Error("failed to join columns as expected")
	}
}
