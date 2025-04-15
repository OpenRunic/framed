package framed

import "fmt"

// defines custom error with details about the table, row and column (if any)
type TableError struct {
	Row     int
	Col     int
	ColName string
	Reason  string
	Source  error
}

func (e TableError) Error() string {
	msg := e.Source.Error()

	if e.Row > -1 {
		msg += fmt.Sprintf(" [row=%d]", e.Row)
	}
	if e.Col > -1 {
		if len(e.ColName) > 0 {
			msg += fmt.Sprintf(" [col=(%d) %s]", e.Col, e.ColName)
		} else {
			msg += fmt.Sprintf(" [col=%d]", e.Col)
		}
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
//	err := NewError($0, $1, $2, $3, $4)
func NewError(err error, reason string) TableError {
	return TableError{
		Reason: reason,
		Row:    -1,
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
//	err := RowError($0, $1, $2, $3, $4)
func RowError(idx int, err error, reason string) TableError {
	return TableError{
		Row:    idx,
		Source: err,
		Reason: reason,
	}
}

// ColError will create an error instance with info about row
// and column index and name(if provided) along with the reason.
//
//	$0 : index of row
//	$1 : index of column
//	$2 : name of column
//	$3 : main originating error
//	$4 : reason of the error
//
//	err := ColError($0, $1, $2, $3, $4)
func ColError(idx int, idx2 int, col string, err error, reason string) *TableError {
	return &TableError{
		Row:     idx,
		Col:     idx2,
		ColName: col,
		Reason:  reason,
		Source:  err,
	}
}
