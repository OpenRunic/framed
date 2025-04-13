package framed

import (
	"fmt"
	"reflect"
	"strings"
)

// splits string to slice separated by 'sep' but respects quotes
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

// joins slice to string joined by 'sep' and uses quotes when required
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

// try to read column value from row as defined type
func TryColumnValue[T any](row *Row, idx int) (T, error) {
	isPtr := false
	val := row.At(idx)
	tp := reflect.TypeOf(val)
	if tp.Kind() == reflect.Ptr {
		isPtr = true
	}

	sample := ToType[*T](nil).Elem()
	if isPtr {
		sample = ToType[*T](nil)
	}

	if tp != sample {
		var v any = ""
		return v.(T), ColError(
			row.Index, idx, "",
			fmt.Errorf("get column value failed; %s != %s", tp, sample),
			"read_value",
		)
	}

	return val.(T), nil
}

// read column value with fallback value
func ColumnValue[T any](row *Row, idx int, def T) T {
	val, err := TryColumnValue[T](row, idx)
	if err != nil {
		return def
	}

	return val
}
