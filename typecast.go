package framed

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// BoolToStringValues contains slice of string
// values that can be auto-translated to boolean
var BoolToStringValues = []string{"1", "false", "no", "n"}

// BoolTrueStringValues contains slice of string
// value that can be auto-translated to true boolean
var BoolTrueStringValues = []string{"1", "true", "yes", "y"}

// StringTranslatableKinds contains slice of [reflect.Kind]
// that can be translated to string
var StringTranslatableKinds = []reflect.Kind{
	reflect.String,
	reflect.Int, reflect.Int32, reflect.Int64,
	reflect.Float32, reflect.Float64,
	reflect.Bool,
}

func init() {
	copy(BoolToStringValues, BoolTrueStringValues)
}

// ToType reads the [reflect.Type] from provided value
func ToType[T any](v T) reflect.Type {
	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Slice {
		return tp.Elem()
	}
	return tp
}

// DetectValueType detects and returns [reflect.Type] of string value
func DetectValueType(name string, value string) reflect.Type {
	sLen := len(name)
	sName := strings.ToLower(name)
	if slices.Contains([]string{"id"}, sName) || (sLen > 3 && name[len(sName)-3:] == "_id") {
		return ToType(0)
	}

	if slices.Contains(BoolToStringValues, strings.ToLower(value)) {
		return ToType(false)
	}

	matched, _ := regexp.MatchString("^[0-9]+$", value)
	if matched {
		_, err := strconv.Atoi(value)
		if err == nil {
			return ToType[int64](0)
		}
	}

	matched, _ = regexp.MatchString("^[0-9.]+$", value)
	if matched {
		_, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return ToType(0.1)
		}
	}

	return ToType("")
}

// ParseInt converts string to int as base10 with variable bit size
func ParseInt(s string, bitSize int) (int64, error) {
	val, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// ConvertValueType converts string to provided type or throws error
func ConvertValueType(value string, tp reflect.Type) (any, error) {
	kind := tp.Kind()
	switch kind {
	case reflect.Bool:
		return slices.Contains(BoolTrueStringValues, strings.ToLower(value)), nil

	case reflect.Int:
		val, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return val, nil

	case reflect.Int32:
		val, err := ParseInt(value, 32)
		if err != nil {
			return nil, err
		}
		return int32(val), nil

	case reflect.Int64:
		val, err := ParseInt(value, 64)
		if err != nil {
			return nil, err
		}
		return val, nil

	case reflect.Float32:
		val, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, err
		}

		return float32(val), nil

	case reflect.Float64:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	if len(value) < 1 {
		return nil, nil
	}

	return value, nil
}

// ColumnValueEncoder encodes the column value to string
// using definition or with default methods
func ColumnValueEncoder(def *Definition, value any) (string, error) {
	if def.Encoder != nil {
		return def.Encoder(value)
	}

	if def.Kind() == reflect.String || slices.Contains(StringTranslatableKinds, def.Kind()) {
		return fmt.Sprintf("%v", value), nil
	}

	return "", fmt.Errorf("failed to encode %s value to string", def.Type)
}

// ColumnValueDecoder decodes the column value from string
// using column definition or with default methods
func ColumnValueDecoder(def *Definition, value string) (any, error) {
	if def.Decoder != nil {
		return def.Decoder(value)
	}

	if def.Kind() == reflect.String {
		return value, nil
	} else if slices.Contains(StringTranslatableKinds, def.Kind()) {
		val, err := ConvertValueType(value, def.Type)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return nil, fmt.Errorf("failed to decode value to %s", def.Type)
}
