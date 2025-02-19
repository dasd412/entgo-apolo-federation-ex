package version

import "fmt"

type InvalidVersionError struct {
	Message string
}

func (e *InvalidVersionError) Error() string {
	return e.Message
}

func NewInvalidVersionError(version string) *InvalidVersionError {
	return &InvalidVersionError{
		Message: fmt.Sprintf("invalid semantic version: %s", version),
	}
}
