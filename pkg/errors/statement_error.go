package errors

import "fmt"

/////////////////////
// Initialization  //
/////////////////////
func NewStatementError(err error, msg string) *StatementError {
	return &StatementError{original: err, message: msg}
}

/////////////////////
// Messages        //
/////////////////////
func StatementPreparationErrorMessage(queryName string) string {
	return fmt.Sprintf("error preparing statement: %s", queryName)
}

func StatementNotPreparedErrorMessage(queryName string) string {
	return fmt.Sprintf("prepared statement %s not found", queryName)
}

/////////////////////
// Structs         //
/////////////////////
type StatementError struct {
	original error
	message  string
}

func (e *StatementError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %s", e.message, e.original.Error())
	}

	return e.message
}
