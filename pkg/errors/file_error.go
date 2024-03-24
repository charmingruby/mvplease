package errors

import "fmt"

// ///////////////////
// Initialization  //
// ///////////////////
func NewInvalidMimetypeError(err error, mimetypeUnmatched string, validMimetypes []string) *FileError {
	return &FileError{
		original: err,
		message:  invalidMimetypeErrorMessage(mimetypeUnmatched, validMimetypes),
	}
}

func NewExceedsMaxSizeError(err error, currentSize, maxSize int64) *FileError {
	return &FileError{
		original: err,
		message:  exceedsMaxiumFileSizeErrorMessage(currentSize, maxSize),
	}
}

func NewEmptyFileError(err error, key string) *FileError {
	return &FileError{
		original: err,
		message:  noFileErrorMessage(key),
	}
}

// ///////////////////
// Messages        //
// ///////////////////
func invalidMimetypeErrorMessage(mimetypeUnmatched string, validMimetypes []string) string {
	var msg = fmt.Sprintf(".%s is not a valid mimetype, please provide a valid mimetype: ", mimetypeUnmatched)

	for idx, mimetype := range validMimetypes {
		if idx+1 == len(mimetypeUnmatched) {
			msg += fmt.Sprintf("or .%s", mimetype)
			continue
		}

		msg += fmt.Sprintf(".%s, ", mimetype)
	}

	return msg
}

func exceedsMaxiumFileSizeErrorMessage(currentSize, maxSize int64) string {
	return fmt.Sprintf("%d perpasses the limit of %d bytes", currentSize, maxSize)
}

func noFileErrorMessage(key string) string {
	return fmt.Sprintf("no file found for key '%s'", key)
}

// ///////////////////
// Structs         //
// ///////////////////
type FileError struct {
	original error
	message  string
}

func (e *FileError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %s", e.message, e.original.Error())
	}

	return e.message
}
