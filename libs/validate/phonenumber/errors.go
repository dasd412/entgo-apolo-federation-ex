package phone_number

import "fmt"

type InvalidPhoneNumberError struct {
	Message string
}

func (e *InvalidPhoneNumberError) Error() string {
	return e.Message
}

func NewInvalidPhoneNumberError(phoneNumber string) *InvalidPhoneNumberError {
	return &InvalidPhoneNumberError{
		Message: fmt.Sprintf("invalid phone number format : %s", phoneNumber),
	}
}
