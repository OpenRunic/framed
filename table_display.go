package framed

import (
	"bytes"
	"fmt"
)

func (t *Table) String() string {
	rows := t.Length()
	sampleSize := t.Options.SampleSize
	previewSize := min(t.Options.Sampling, rows)

	buf := bytes.NewBufferString("\n")

	buf.WriteString(fmt.Sprintf("\nColumns (%d):\n------------", t.ColLength()))
	for idx, col := range t.State.Columns {
		def := t.State.Definition(col)
		samples := make([]string, previewSize)
		for sI := range previewSize {
			sVal, err := ColumnValueEncoder(def, t.At(sI).At(idx))
			if err == nil {
				if len(sVal) > sampleSize {
					samples[sI] = sVal[0:sampleSize] + "..."
				} else {
					samples[sI] = sVal
				}
			}
		}

		samplesStr := ""
		if previewSize > 0 {
			samplesStr = fmt.Sprintf(" = %s", JoinAtChar(samples, t.Options.Separator))
		}

		buf.WriteString(fmt.Sprintf("\n#%d %s (%s)%s", idx, col, t.State.DataType(col), samplesStr))
	}

	buf.WriteString(fmt.Sprintf("\n\nTotal rows: %d\n", rows))

	return buf.String()
}
