package image

import "fmt"

type InvalidImageUrlError struct {
	Message string
}

func (e *InvalidImageUrlError) Error() string {
	return e.Message
}

func NewInvalidImageUrlError(imageUrl string) *InvalidImageUrlError {
	return &InvalidImageUrlError{
		Message: fmt.Sprintf("invalid image URL format : %s", imageUrl),
	}
}
