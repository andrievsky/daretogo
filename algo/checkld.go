package algo

import (
	"unicode"
)

// LetterAndDigitCases test strings
var LetterAndDigitCases = [...]string{"abc", "123", "abc123"}

// LetterAndDigitResults test results
var LetterAndDigitResults = [...]bool{false, false, true}

// CheckForLetterAndDigit check a string has at least one letter and digit
func CheckForLetterAndDigit(s string) bool {
	var hasLetter = false
	var hasDigit = false
	for _, r := range s {
		if !unicode.IsLetter(r) {
			hasLetter = true
		}
		if !unicode.IsDigit(r) {
			hasDigit = true
		}
		if hasLetter && hasDigit {
			return true
		}
	}
	return false
}

// TODO Update to a valid regex
/*
// CheckForLetterAndDigitWithRegex same check with Regex
func CheckForLetterAndDigitWithRegex(s string) bool {
	var validID = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*[0-9]).+$`)
	return validID.MatchString(s)
}
*/
