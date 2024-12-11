package Validator

import (
	"regexp"
	"testing"
)

var emailRegexp = regexp.MustCompile("(?i)" + "^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + "@" + "[a-z0-9-]+(\\.[a-z0-9-]+)*$")
/*
ROOST_METHOD_HASH=IsValidEmail_ea24af8bd9
ROOST_METHOD_SIG_HASH=IsValidEmail_d2603fb18f


 */
func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}

func TestIsValidEmail(t *testing.T) {

	type testCase struct {
		description string
		email       string
		expected    bool
	}

	testCases := []testCase{
		{
			description: "Maximum Email Length Limitation",

			email:    "a" + string(make([]byte, 243)) + "@example.com",
			expected: false,
		},
		{
			description: "Validate Email With Special Characters",
			email:       "user.name+tag+sorting@example.com",
			expected:    true,
		},
		{
			description: "Validate Email Domains With Subdomains",
			email:       "email@subdomain.example.com",
			expected:    true,
		},
		{
			description: "Check TLD and Domain Edge Cases",
			email:       "email@domain.invalid",
			expected:    false,
		},
		{
			description: "Test Case Sensitivity of Local Part",
			email:       "User.Name@example.com",
			expected:    true,
		},
		{
			description: "Validate Empty String as Email",
			email:       "",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.description)
			actual := IsValidEmail(tc.email)
			if actual != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.description, tc.expected, actual)
			} else {
				t.Logf("Passed %s", tc.description)
			}
		})
	}
}

