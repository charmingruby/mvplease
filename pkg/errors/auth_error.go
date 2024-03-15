package errors

/////////////////////
// Initialization  //
/////////////////////
func NewInvalidCredentialsError() *InvalidCredentialsError {
	return &InvalidCredentialsError{
		message: "Invalid credentials.",
	}
}

/////////////////////
// Messages        //
/////////////////////

/////////////////////
// Structs         //
/////////////////////
type InvalidCredentialsError struct {
	message string
}

func (e *InvalidCredentialsError) Error() string {
	return e.message
}
