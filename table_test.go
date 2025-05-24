package framed_test

import (
	"math/rand"
	"strings"
	"time"

	"fmt"

	"github.com/OpenRunic/framed"
)

const stringSet = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

type TestAddressInfo struct {
	Street string
	City   string
}

func StringRandom(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = stringSet[seededRand.Intn(len(stringSet))]
	}
	return string(b)
}

func SampleTestTable(cbs ...framed.OptionCallback) *framed.Table {
	finalCbs := make([]framed.OptionCallback, 0)
	finalCbs = append(finalCbs,
		framed.WithColumns(
			"id", "first_name", "last_name", "age", "address", "phone",
		),
		framed.WithDefinition(
			"address",
			framed.NewDefinition(framed.ToType[*TestAddressInfo](nil)).
				UseEncoder(func(a any) (string, error) {
					addr := a.(*TestAddressInfo)
					return fmt.Sprintf("%s, %s", addr.Street, addr.City), nil
				}).
				UseDecoder(func(s string) (any, error) {
					splits := strings.Split(s, ",")
					return &TestAddressInfo{
						Street: strings.Trim(splits[0], " "),
						City:   strings.Trim(splits[1], " "),
					}, nil
				}),
		),
	)
	finalCbs = append(finalCbs, cbs...)

	size := 10000
	phoneMin := 1000000000
	rows := make([][]string, size)
	for idx := range size {
		rows[idx] = []string{
			fmt.Sprint(idx + 1),
			StringRandom(seededRand.Intn(3) + 4),
			StringRandom(seededRand.Intn(7) + 5),
			fmt.Sprint(seededRand.Intn(45) + 5),
			StringRandom(10) + ", " + StringRandom(5),
			fmt.Sprint(seededRand.Intn(9999999999-phoneMin) + phoneMin),
		}
	}

	df, _ := framed.Raw(rows,
		finalCbs...,
	)

	return df
}
