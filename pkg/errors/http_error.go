package errors

import "fmt"

/////////////////////
// Initialization  //
/////////////////////
func NewPayloadError(err error, msg string) *PayloadError {
	return &PayloadError{original: err, message: msg}
}

/////////////////////
// Messages        //
/////////////////////
func PayloadErrorMessage() string {
	return "Payload error"
}

func TokenRetrievingErrorMessage() string {
	return "Cannot retrieve payload from token"
}

/////////////////////
// Structs         //
/////////////////////
type PayloadError struct {
	original error
	message  string
}

func (e *PayloadError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %s", e.message, e.original.Error())
	}

	return e.message
}
