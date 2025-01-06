package middleware

import (
	"os"
	"testing"
	"github.com/joho/godotenv"
)


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestgoDotEnvVariable(t *testing.T) {

	type testCase struct {
		name           string
		envFileContent string
		key            string
		expectedValue  string
		shouldError    bool
	}

	tests := []testCase{
		{
			name:           "Successfully Retrieve an Environment Variable",
			envFileContent: "KEY=VALUE\n",
			key:            "KEY",
			expectedValue:  "VALUE",
			shouldError:    false,
		},
		{
			name:           "Handle Missing Environment Variable Gracefully",
			envFileContent: "",
			key:            "NON_EXISTENT_KEY",
			expectedValue:  "",
			shouldError:    false,
		},
		{
			name:           "Handle Error Loading .env File",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  "",
			shouldError:    false,
		},
		{
			name:           "Test with Multiple Environment Variables",
			envFileContent: "KEY1=VALUE1\nKEY2=VALUE2\n",
			key:            "KEY1",
			expectedValue:  "VALUE1",
			shouldError:    false,
		},
		{
			name:           "Check for Environment Variable with Whitespace",
			envFileContent: "KEY= VALUE \n",
			key:            "KEY",
			expectedValue:  " VALUE ",
			shouldError:    false,
		},
		{
			name:           "Environment Variable with Special Characters",
			envFileContent: "KEY=Value#1$\n",
			key:            "KEY",
			expectedValue:  "Value#1$",
			shouldError:    false,
		},
		{
			name:           "Verify Behavior with an Unloaded .env File",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  os.Getenv("KEY"),
			shouldError:    false,
		},
		{
			name:           "Check Error Logging on .env Loading Failure",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  "",
			shouldError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.envFileContent != "" {
				tmpFile, err := os.CreateTemp("", ".env")
				if err != nil {
					t.Fatalf("could not create temp file: %v", err)
				}
				defer os.Remove(tmpFile.Name())
				if _, err := tmpFile.WriteString(tt.envFileContent); err != nil {
					t.Fatalf("could not write to temp file: %v", err)
				}
				if err := tmpFile.Close(); err != nil {
					t.Fatalf("could not close temp file: %v", err)
				}

				if err := godotenv.Load(tmpFile.Name()); err != nil {
					if !tt.shouldError {
						t.Errorf("failed to load .env file: %v", err)
					}
				}
			} else {

				os.Remove(".env")
			}

			result := goDotEnvVariable(tt.key)

			if result != tt.expectedValue {
				t.Errorf("expected %q, got %q", tt.expectedValue, result)
			}

			t.Logf("Test %s completed: expected %s, got %s", tt.name, tt.expectedValue, result)
		})
	}
}
