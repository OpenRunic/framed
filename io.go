package framed

import (
	"bufio"
	"io"
	"iter"
)

// convert data from reader to iterable bytes
func ReaderToLines(r io.Reader) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		sc := bufio.NewScanner(r)
		if err := sc.Err(); err != nil {
			return
		}

		for sc.Scan() {
			if !yield(sc.Bytes()) {
				return
			}
		}
	}
}
