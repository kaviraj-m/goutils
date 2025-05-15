package validation

import (
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

// IsEmail validates if the string is a valid email address
func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsURL validates if the string is a valid URL
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// IsNumeric checks if a string contains only numeric characters
func IsNumeric(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(str) > 0
}

// IsAlpha checks if a string contains only alphabetic characters
func IsAlpha(str string) bool {
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return len(str) > 0
}

// IsAlphanumeric checks if a string contains only alphanumeric characters
func IsAlphanumeric(str string) bool {
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	return len(str) > 0
}

// MinLength validates the minimum length of a string
func MinLength(str string, min int) bool {
	return len(str) >= min
}

// MaxLength validates the maximum length of a string
func MaxLength(str string, max int) bool {
	return len(str) <= max
}

// IsPhoneNumber validates common phone number formats
// This is a simple implementation, production code might need more sophisticated validation
func IsPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(strings.ReplaceAll(phone, " ", ""))
}
