package middleware

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func TestGoDotEnvVariable_793184d000(t *testing.T) {
	var key, value string

	// Test case 1: Key exists in the .env file
	key = "EXISTING_KEY" // TODO: Replace with a key that exists in your .env file
	value = goDotEnvVariable(key)
	if value == "" {
		t.Error("TestGoDotEnvVariable_793184d000 failed - ", key, "does not exist in .env file")
	} else {
		t.Log("TestGoDotEnvVariable_793184d000 passed - ", key, "exists in .env file")
	}

	// Test case 2: Key does not exist in the .env file
	key = "NON_EXISTING_KEY" // TODO: Replace with a key that does not exist in your .env file
	value = goDotEnvVariable(key)
	if value != "" {
		t.Error("TestGoDotEnvVariable_793184d000 failed - ", key, "exists in .env file")
	} else {
		t.Log("TestGoDotEnvVariable_793184d000 passed - ", key, "does not exist in .env file")
	}
}
