package phone_number

import (
	"regexp"
)

const PhoneNumberPattern = `^010-[0-9]{4}-[0-9]{4}$`

func ValidatePhoneNumber(phoneNumber string) error {
	match, _ := regexp.MatchString(PhoneNumberPattern, phoneNumber)
	if match {
		return nil
	} else {
		return NewInvalidPhoneNumberError(phoneNumber)
	}
}
