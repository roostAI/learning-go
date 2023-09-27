// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package LevenshteinDistance

import (
	"testing"
)

func TestMin_296d0913dd(t *testing.T) {
	var a, b uint16

	// Test case 1: When a is less than b
	a, b = 5, 10
	expected := a
	actual := min(a, b)
	if actual != expected {
		t.Errorf("Test case 1 failed: a(%v), b(%v), expected(%v), returned(%v)", a, b, expected, actual)
	} else {
		t.Logf("Test case 1 succeeded")
	}

	// Test case 2: When a is greater than b
	a, b = 15, 10
	expected = b
	actual = min(a, b)
	if actual != expected {
		t.Errorf("Test case 2 failed: a(%v), b(%v), expected(%v), returned(%v)", a, b, expected, actual)
	} else {
		t.Logf("Test case 2 succeeded")
	}

	// Test case 3: When a is equal to b
	a, b = 10, 10
	expected = a
	actual = min(a, b)
	if actual != expected {
		t.Errorf("Test case 3 failed: a(%v), b(%v), expected(%v), returned(%v)", a, b, expected, actual)
	} else {
		t.Logf("Test case 3 succeeded")
	}
}
