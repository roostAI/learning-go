// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package LevenshteinDistance

import (
	"testing"
)

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func TestLevenshteinDistance_c58c51e350(t *testing.T) {

	t.Run("Test with equal strings", func(t *testing.T) {
		a := "test"
		b := "test"
		expected := 0

		if result := levenshteinDistance(a, b); result != expected {
			t.Errorf("Expected %d, but got %d for inputs %s and %s", expected, result, a, b)
		} else {
			t.Logf("Success: Expected %d, got %d for inputs %s and %s", expected, result, a, b)
		}
	})

	t.Run("Test with one empty string", func(t *testing.T) {
		a := ""
		b := "test"
		expected := len(b)

		if result := levenshteinDistance(a, b); result != expected {
			t.Errorf("Expected %d, but got %d for inputs %s and %s", expected, result, a, b)
		} else {
			t.Logf("Success: Expected %d, got %d for inputs %s and %s", expected, result, a, b)
		}
	})

	t.Run("Test with two different strings", func(t *testing.T) {
		a := "kitten"
		b := "sitting"
		expected := 3

		if result := levenshteinDistance(a, b); result != expected {
			t.Errorf("Expected %d, but got %d for inputs %s and %s", expected, result, a, b)
		} else {
			t.Logf("Success: Expected %d, got %d for inputs %s and %s", expected, result, a, b)
		}
	})
}