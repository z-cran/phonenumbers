package usphonenumbers

import (
	"fmt"
	"regexp"
)

type PhoneNumber struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
}

var (
	// Strips out non-numeric characters when parsing phone numbers
	numericRegex = regexp.MustCompile(`[^0-9]+`)
)

const (
	ErrInvalidNumberLength = "Invalid number length. Accepted lengths: 10, 11."
	ErrInvalidCountryCode  = "Invalid country code. Accepted value: 1."
	ErrInvalidAreaCode     = "Invalid area code."
)

// Parse returns a PhoneNumber if a provided string is a valid US phone number and an error if it is invalid
func Parse(s string) (*PhoneNumber, error) {
	// Strip non-numerical characters
	phone := numericRegex.ReplaceAllString(s, "")

	// Check for valid length:
	//   10 for standard US phone numbers including area codes
	//   11 if country code (1) is included
	length := len(phone)
	if length != 10 && length != 11 {
		return nil, fmt.Errorf("%s: %s", ErrInvalidNumberLength, s)
	}

	// If length is 11, assert first character is 1
	if length == 11 {
		if phone[0:1] != USCountryCode {
			return nil, fmt.Errorf("%s: %s", ErrInvalidCountryCode, s)
		}
		// Strip country code
		phone = phone[1:11]
	}

	// Check for valid area code
	areaCode := phone[0:3]
	var validAreaCode bool
	for _, a := range USAreaCodes {
		if areaCode == a {
			validAreaCode = true
		}
	}
	if !validAreaCode {
		return nil, fmt.Errorf("%s: %s", ErrInvalidAreaCode, s)
	}

	return &PhoneNumber{
		PhoneNumber: phone,
		AreaCode:    areaCode,
	}, nil
}
