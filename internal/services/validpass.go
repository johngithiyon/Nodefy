package services

import "strings"

func IsValidPassword(password string) bool {

	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	specialChars := "@$!%*?&"

	for _, ch := range password {

		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		} else if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		} else if strings.ContainsRune(specialChars, ch) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}