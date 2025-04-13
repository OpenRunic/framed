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

// makes default table error info
func NewError(err error, reason string) TableError {
	return TableError{
		Reason: reason,
		Row:    -1,
		Source: err,
	}
}

// makes row error info
func RowError(idx int, err error, reason string) TableError {
	return TableError{
		Row:    idx,
		Source: err,
		Reason: reason,
	}
}

// makes column error info
func ColError(idx int, idx2 int, col string, err error, reason string) *TableError {
	return &TableError{
		Row:     idx,
		Col:     idx2,
		ColName: col,
		Reason:  reason,
		Source:  err,
	}
}
