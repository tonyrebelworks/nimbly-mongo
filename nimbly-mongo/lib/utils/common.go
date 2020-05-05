package utils

import (
	"regexp"
)

// Contains find is string in slices
func Contains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}

// IsEmail check email the real email string
func IsEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}
