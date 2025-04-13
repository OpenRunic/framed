package framed

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var BoolToStringValues = []string{}
var BoolTrueStringValues = []string{"1", "true", "yes", "y"}
var BoolFalseStringValues = []string{"1", "false", "no", "n"}

func init() {
	// copy data to build final set of boolable values
	copy(BoolToStringValues, BoolTrueStringValues)
	copy(BoolToStringValues, BoolFalseStringValues)
}

// detect the type of value
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
			return ToType[int32](0)
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

// parse int value
func ParseInt(s string, bitSize int) (int64, error) {
	val, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// convert the value to provided type (if possible)
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

// encode the column value to string using column definition
func ColumnValueEncoder(def *ColumnDefinition, value any) (string, error) {
	if def.Encoder != nil {
		return def.Encoder(value)
	}

	if def.Kind() == reflect.String || slices.Contains(StringTranslatableKinds, def.Kind()) {
		return fmt.Sprintf("%v", value), nil
	}

	return "", fmt.Errorf("failed to encode %s value to string", def.Type)
}

// decode the column value from string using column definition
func ColumnValueDecoder(def *ColumnDefinition, value string) (any, error) {
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
