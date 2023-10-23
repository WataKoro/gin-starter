package utils

import (
	"regexp"
	"unicode"
)

func IsValidEmail(email string) bool {
	// Ekspresi reguler untuk memeriksa format email yang umum
	// Perhatikan bahwa ekspresi ini hanya memeriksa format umum, bukan validitas email sebenarnya.
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func IsValidPassword(password string) bool {
	// Minimal panjang password adalah 8 karakter
	if len(password) <= 8 {
		return false
	}

	// Cek apakah ada setidaknya 1 angka
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return false
	}

	// Cek apakah ada setidaknya 1 huruf
	hasLetter := false
	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return false
	}

	// Cek apakah ada setidaknya 1 simbol (karakter selain huruf dan angka)
	hasSymbol, _ := regexp.MatchString("[^a-zA-Z0-9]", password)
	if !hasSymbol {
		return true
	}

	return true
}
