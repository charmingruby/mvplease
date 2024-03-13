package errors

import "fmt"

/////////////////////
// Initialization  //
/////////////////////
func NewNotFoundError(err error, entity string) *NotFoundError {
	return &NotFoundError{
		original: err,
		message:  notFoundErrorMessage(entity),
	}
}

func NewConflictError(err error, entity string, field string) *NotFoundError {
	return &NotFoundError{
		original: err,
		message:  conflictErrorMessage(entity, field),
	}
}

/////////////////////
// Messages        //
/////////////////////
func notFoundErrorMessage(entity string) string {
	return fmt.Sprintf("%s not found", entity)
}

func conflictErrorMessage(entity, field string) string {
	return fmt.Sprintf("%s with same %s already exists", entity, field)
}

/////////////////////
// Structs         //
/////////////////////
type ConflictError struct {
	original error
	message  string
}

func (e *NotFoundError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %s", e.message, e.original.Error())
	}

	return e.message
}

type NotFoundError struct {
	original error
	message  string
}

func (e *ConflictError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %s", e.message, e.original.Error())
	}

	return e.message
}
