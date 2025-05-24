package framed

import (
	"fmt"
	"reflect"
	"strings"
)

// SplitAtChar splits string to slice separated
// by provided separator but respects quotes
//
//	$0 : string to split
//	$1 : separator to split the string
//
//	splits := SplitAtChar($0, $1)
func SplitAtChar(s string, sep byte) []string {
	var beg int
	var inString bool
	res := []string{}

	for i := range len(s) {
		if s[i] == sep && !inString {
			res = append(res, strings.Trim(s[beg:i], "\""))
			beg = i + 1
		} else if s[i] == '"' {
			if !inString {
				inString = true
			} else if i > 0 && s[i-1] != '\\' {
				inString = false
			}
		}
	}

	return append(res, strings.Trim(s[beg:], "\""))
}

// JoinAtChar joins slice of string to string joined
// by provided separator and add quotes when needed
//
//	$0 : slice of String to join
//	$1 : joining to split the string
//
//	splits := JoinAtChar($0, $1)
func JoinAtChar(ss []string, sep byte) string {
	res := ""

	sepStr := string(sep)
	for i, s := range ss {
		if i > 0 {
			res += sepStr
		}

		if strings.Contains(s, sepStr) {
			res += fmt.Sprintf("\"%s\"", s)
		} else {
			res += s
		}
	}

	return res
}

// TryColumnValue will try to read the value of column as
// provided data type or else create error message.
//
//	$0 : row instance from table
//	$1 : index of the column
//
//	val, err := TryColumnValue[T any]($0, $1)
func TryColumnValue[T any](row *Row, idx int) (T, error) {
	isPtr := false
	val := row.At(idx)
	tp := reflect.TypeOf(val)
	if tp == nil {
		return any("").(T), ColumnReadValueNilError(row.Index, idx)
	}
	if tp.Kind() == reflect.Ptr {
		isPtr = true
	}

	sample := ToType[*T](nil).Elem()
	if isPtr {
		sample = ToType[*T](nil)
	}

	if tp != sample {
		return any("").(T), ColumnReadValueError(row.Index, idx, tp, sample)
	}

	return val.(T), nil
}

// ColumnValue will try to read the value of column as
// provided data type or else returns fallback value.
//
//	$0 : row instance from table
//	$1 : index of the column
//	$2 : fallback value of the type
//
//	val, err := ColumnValue[T any]($0, $1, $2)
func ColumnValue[T any](row *Row, idx int, def T) T {
	val, err := TryColumnValue[T](row, idx)
	if err != nil {
		return def
	}

	return val
}
