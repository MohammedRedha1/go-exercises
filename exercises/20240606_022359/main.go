package main

import "unicode"

func isValidPassword(password string) bool {
	var charsCount, uppersCount, digitsCount int = 0, 0, 0
	for _, char := range password {
		charsCount++
		if charsCount > 12 {
			return false
		}
		if unicode.IsUpper(char) {
			uppersCount++
		}
		if unicode.IsDigit(char) {
			digitsCount++
		}

	}
	if uppersCount >= 1 && digitsCount >= 1 {
		return true
	}
	return false
}
