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

func TestMin_296d0913dd(t *testing.T) {
	testCases := []struct {
		name string
		a    uint16
		b    uint16
		want uint16
	}{
		{
			name: "Test Case 1: a is less than b",
			a:    5,
			b:    10,
			want: 5,
		},
		{
			name: "Test Case 2: a is greater than b",
			a:    15,
			b:    10,
			want: 10,
		},
		{
			name: "Test Case 3: a is equal to b",
			a:    10,
			b:    10,
			want: 10,
		},
		{
			name: "Test Case 4: a and b are zero",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := min(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("min(%v, %v): got %v, want %v", tc.a, tc.b, got, tc.want)
			} else {
				t.Logf("Success: min(%v, %v): got %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
