package framed

import (
	"errors"
	"fmt"
	"reflect"
)

// defines custom error with details about the table, row and column (if any)
type TableError struct {
	Row    int
	Col    int
	Reason string
	Source error
}

func (e TableError) Error() string {
	msg := e.Source.Error()

	if e.Row > -1 {
		msg += fmt.Sprintf(" [row=%d]", e.Row)
	}
	if e.Col > -1 {
		msg += fmt.Sprintf(" [col=%d]", e.Col)
	}
	if len(e.Reason) > 0 {
		msg += fmt.Sprintf(" [reason=%s]", e.Reason)
	}

	return msg
}

// NewError will create an error instance along with the reason.
//
//	$0 : main originating error
//	$1 : reason of the error
//
//	err := NewError($0, $1)
func NewError(err error, reason string) TableError {
	return TableError{
		Reason: reason,
		Row:    -1,
		Col:    -1,
		Source: err,
	}
}

// RowError will create an error instance with info about
// row along with the reason.
//
//	$0 : index of row
//	$1 : main originating error
//	$2 : reason of the error
//
//	err := RowError($0, $1, $2)
func RowError(idx int, err error, reason string) TableError {
	return TableError{
		Row:    idx,
		Col:    -1,
		Source: err,
		Reason: reason,
	}
}

// ColError will create an error instance with info about row
// and column index and name(if provided) along with the reason.
//
//	$0 : index of row
//	$1 : index of column
//	$2 : main originating error
//	$3 : reason of the error
//
//	err := ColError($0, $1, $2, $3)
func ColError(idx int, idx2 int, err error, reason string) TableError {
	return TableError{
		Row:    idx,
		Col:    idx2,
		Reason: reason,
		Source: err,
	}
}

// InvalidColumnValueTypeError creates error for non-matching value type
func InvalidTableShapeError(x int, y int) TableError {
	return NewError(
		fmt.Errorf("mismatching size of columns and values; %d != %d", x, y),
		"invalid_shape",
	)
}

// InvalidColumnValueTypeError creates error for non-matching value type
func InvalidColumnValueTypeError(column string, eType reflect.Type) TableError {
	return NewError(fmt.Errorf("non-matching type of value [%s]; expected %s", column, eType), "invalid_type")
}

// NonNumericColumnValueError creates error when column value cannot be read as numberic
func NonNumericColumnValueError(column string) TableError {
	return NewError(fmt.Errorf("unable to resolve numeric value from column: %s", column), "non_numeric")
}

// WriteColumnValueError creates error for failing to write column value
func WriteColumnValueError(idx int, idx2 int, rType reflect.Type, sType reflect.Type) TableError {
	return ColError(
		idx, idx2,
		fmt.Errorf("set column value failed; %s != %s", rType, sType),
		"write_value",
	)
}

// EncodeColumnValueError creates error for failing to encode column value to string
func EncodeColumnValueError(idx int, idx2 int, err error) TableError {
	return ColError(idx, idx2, err, "value_encode")
}

// DecodeColumnValueError creates error for failing to decode strng to column value
func DecodeColumnValueError(idx int, idx2 int, err error) TableError {
	return ColError(idx, idx2, err, "value_decode")
}

// UnknownColumnError creates error for missing column action
func UnknownColumnError(column string) TableError {
	return NewError(fmt.Errorf("unable to locate column: %s", column), "unknown_column")
}

// FailedToResolveColumnDataError creates error for failing to resolve column data
func FailedToResolveColumnDataError(column string) TableError {
	return NewError(fmt.Errorf("unable to load data for column: %s", column), "column_data_missing")
}

// ColumnReadValueNilError creates error when read value is nil
func ColumnReadValueNilError(idx int, idx2 int) TableError {
	return ColError(
		idx, idx2,
		errors.New("get column value failed; found nil"),
		"read_value_nil",
	)
}

// ColumnReadValueError creates error for inability to read column value
func ColumnReadValueError(idx int, idx2 int, rType reflect.Type, sType reflect.Type) TableError {
	return ColError(
		idx, idx2,
		fmt.Errorf("get column value failed; %s != %s", rType, sType),
		"read_value",
	)
}

// RowUnknownColumnError creates error for missing column action
func RowUnknownColumnError(index int, column string) TableError {
	return RowError(index, fmt.Errorf("unable to locate column: %s", column), "unknown_column")
}

// RowUnknownColumnIndexError creates error for missing column action
func RowUnknownColumnIndexError(index int, index2 int) TableError {
	return RowError(index, fmt.Errorf("unable to locate column at %d", index2), "unknown_column")
}

// RowValidationFailedError creates error for validation
func RowValidationFailedError(index int, err error) TableError {
	return RowError(index, err, "validation_failed")
}
