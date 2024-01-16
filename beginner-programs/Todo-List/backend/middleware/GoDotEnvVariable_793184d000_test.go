package middleware

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func TestGoDotEnvVariable_793184d000(t *testing.T) {
	// Test case 1: when .env file is loaded successfully
	err := godotenv.Load(".env")
	if err != nil {
		t.Error("Failed to load .env file in test setup")
	}
	os.Setenv("TEST_KEY", "TEST_VALUE")
	value := goDotEnvVariable("TEST_KEY")
	assert.Equal(t, "TEST_VALUE", value, "Expected TEST_VALUE, but got ", value)

	// Test case 2: when .env file is not present
	err = os.Remove(".env")
	if err != nil {
		t.Error("Failed to remove .env file in test setup")
	}
	value = goDotEnvVariable("TEST_KEY")
	assert.Equal(t, "", value, "Expected empty string, but got ", value)

	// Test case 3: when key is not present in the .env file
	err = godotenv.Load(".env")
	if err != nil {
		t.Error("Failed to load .env file in test setup")
	}
	value = goDotEnvVariable("NON_EXISTENT_KEY")
	assert.Equal(t, "", value, "Expected empty string, but got ", value)
}
