package framed_test

import (
	"math/rand"
	"time"

	"fmt"

	"github.com/OpenRunic/framed"
)

const stringSet = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringRandom(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = stringSet[seededRand.Intn(len(stringSet))]
	}
	return string(b)
}

func SampleTestTable(sizes ...int) *framed.Table {
	size := 100
	if len(sizes) > 0 {
		size = sizes[0]
	}

	rows := make([][]string, size)
	for idx := range size {
		rows[idx] = []string{
			fmt.Sprint(idx + 1), StringRandom(10), StringRandom(10), fmt.Sprint(rand.Intn(50)),
		}
	}

	df, _ := framed.Raw(rows,
		framed.WithColumns(
			"id", "first_name", "last_name", "age",
		),
	)

	return df
}
