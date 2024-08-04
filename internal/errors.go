package internal

import (
	"errors"
	"fmt"
	"os"
)

// Errors for file handling
var OpenInputDirError = &fileError{msg: "error opening Input directory"}
var OpenOutputDirError = &fileError{msg: "error opening Output directory"}
var OpenFileError = &fileError{msg: os.ErrNotExist.Error()}
var ReadFileError = &fileError{msg: "error reading xml file content"}
var DecodeFileError = &fileError{msg: "error decoding xml file content"}
var ParseFileError = &fileError{msg: "error parsing xml file to json"}

type fileError struct {
	msg string
}

func (e *fileError) Error() string {
	return e.msg
}
func (e *fileError) Unwrap() error {
	return errors.New(e.msg)
}

/**
 * throwError is a helper function to wrap errors with custom errors
 * @param custom error
 * @param err error
 * @return error
 */
func throwError(custom error, err error) error {
	return fmt.Errorf("%w-%v", custom, err)
}

// Path: internal/errors.go
