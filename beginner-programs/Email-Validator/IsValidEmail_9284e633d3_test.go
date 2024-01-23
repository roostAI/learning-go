package Validator

import (
	"regexp"
	"testing"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.+\\\"]+@[a-zA-Z0-9-\\.]+\\.[a-zA-Z]{2,}$")

func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}

func TestIsValidEmail(t *testing.T) {
	// Test case 1: Valid email
	email := "test@example.com"
	if !IsValidEmail(email) {
		t.Errorf("Expected %s to be a valid email", email)
	}

	// Test case 2: Email with length greater than 254
	email = "test@example.com" + "a"
	if IsValidEmail(email) {
		t.Errorf("Expected %s to be an invalid email", email)
	}

	// Test case 3: Email with special characters
	email = "test@example.com"
	if !IsValidEmail(email) {
		t.Errorf("Expected %s to be a valid email", email)
	}

	// Test case 4: Email without @
	email = "test.example.com"
	if IsValidEmail(email) {
		t.Errorf("Expected %s to be an invalid email", email)
	}
}
