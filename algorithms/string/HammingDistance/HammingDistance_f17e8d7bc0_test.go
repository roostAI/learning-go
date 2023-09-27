// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package HammingDistance

import (
	"log"
	"testing"
)

func hammingDistance(a, b string) int {
	if len(a) != len(b) {
		log.Fatal("Strings are of different length")
	}

	var distance int = 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance
}

func TestHammingDistance_f17e8d7bc0(t *testing.T) {
	// Test case 1: Strings of equal length and no differences
	a := "test"
	b := "test"
	expected := 0
	result := hammingDistance(a, b)
	if result != expected {
		t.Errorf("Test case 1 failed, hammingDistance('%s', '%s') returned %d, expected %d", a, b, result, expected)
	} else {
		t.Logf("Test case 1 succeeded")
	}

	// Test case 2: Strings of equal length and some differences
	a = "test"
	b = "tent"
	expected = 1
	result = hammingDistance(a, b)
	if result != expected {
		t.Errorf("Test case 2 failed, hammingDistance('%s', '%s') returned %d, expected %d", a, b, result, expected)
	} else {
		t.Logf("Test case 2 succeeded")
	}

	// Test case 3: Strings of different lengths
	a = "test"
	b = "testing"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test case 3 failed, hammingDistance('%s', '%s') was supposed to panic", a, b)
		} else {
			t.Logf("Test case 3 succeeded")
		}
	}()
	hammingDistance(a, b)
}
