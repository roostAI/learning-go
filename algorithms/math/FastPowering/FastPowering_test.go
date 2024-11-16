// ********RoostGPT********
/*
Test generated by RoostGPT for test unit-golang using AI Type  and AI Model 

ROOST_METHOD_HASH=fastPowering_860ba92ce3
ROOST_METHOD_SIG_HASH=fastPowering_be2e642e07

Existing Test Information:
These test cases are already implemented and not included for test generation scenario:
File: learning-go/algorithms/math/FastPowering/FastPowering_be2e642e07_test.go
Test Cases:
    [TestFastPowering_be2e642e07]

File: learning-go/algorithms/math/FastPowering/fastpowering_test.go
Test Cases:
    [TestFastPowering]

================================VULNERABILITIES================================
Vulnerability: CWE-682: Incorrect Calculation
Issue: The function may produce incorrect results or enter an infinite recursion for large power values due to integer division truncation in power/2
Solution: Use float64(power)/2 instead of power/2 to avoid integer division truncation

Vulnerability: CWE-1333: Inefficient Regular Expression Complexity
Issue: The recursive approach may lead to stack overflow for very large power values, causing a denial of service
Solution: Implement an iterative version of the algorithm using bit manipulation to handle large exponents efficiently

Vulnerability: CWE-190: Integer Overflow or Wraparound
Issue: Potential integer overflow when calculating power/2 for large negative power values
Solution: Add input validation to ensure power is non-negative, or use uint instead of int for the power parameter

================================================================================
Based on the provided function and requirements, here are some test scenarios for the `fastPowering` function:

Scenario 1: Large Positive Base and Power

Details:
  Description: Test the function with a large positive base and power to ensure it can handle larger numbers correctly.
Execution:
  Arrange: Prepare a large base and power.
  Act: Call fastPowering with base 5 and power 20.
  Assert: Verify that the result matches the expected value (95367431640625).
Validation:
  This test ensures the function can handle larger computations accurately, which is crucial for real-world applications that might involve significant calculations.

Scenario 2: Negative Base with Odd Power

Details:
  Description: Verify the function's behavior with a negative base raised to an odd power.
Execution:
  Arrange: Set up a negative base and an odd power.
  Act: Call fastPowering with base -3 and power 5.
  Assert: Check if the result is equal to -243.
Validation:
  This test is important to ensure the function correctly handles negative bases, especially with odd powers where the result should be negative.

Scenario 3: Fractional Base with Positive Power

Details:
  Description: Test the function's ability to handle fractional bases with positive powers.
Execution:
  Arrange: Prepare a fractional base and a positive power.
  Act: Call fastPowering with base 0.5 and power 4.
  Assert: Verify that the result is close to 0.0625 (considering floating-point precision).
Validation:
  This scenario tests the function's accuracy with fractional bases, which is important for applications dealing with non-integer values.

Scenario 4: Very Large Power

Details:
  Description: Evaluate the function's performance and accuracy with a very large power.
Execution:
  Arrange: Set up a moderate base with a very large power.
  Act: Call fastPowering with base 2 and power 1000000.
  Assert: Verify that the result is a very large number (or infinity, depending on float64 limitations).
Validation:
  This test checks the function's ability to handle extreme cases and its behavior near or at the limits of float64 representation.

Scenario 5: Base 1 with Any Power

Details:
  Description: Confirm that any power of 1 always results in 1.
Execution:
  Arrange: Use base 1 with various powers (positive, negative, and zero).
  Act: Call fastPowering with base 1 and powers like 100, -50, and 0.
  Assert: Check that all results are exactly 1.
Validation:
  This test ensures the function correctly handles the special case of base 1, which should always return 1 regardless of the power.

Scenario 6: Very Small Fractional Base with Large Power

Details:
  Description: Test the function's precision with a very small fractional base raised to a large power.
Execution:
  Arrange: Prepare a very small fractional base and a large power.
  Act: Call fastPowering with base 0.1 and power 100.
  Assert: Verify that the result is very close to 1e-100 (considering floating-point precision).
Validation:
  This scenario tests the function's ability to handle very small results accurately, which is crucial for scientific or financial calculations involving small numbers.

These scenarios cover a range of cases including normal operation, edge cases, and potential numerical precision issues. They test the function's behavior with various types of inputs, ensuring comprehensive coverage of its functionality.
*/

// ********RoostGPT********
package fastpowering

import (
	"math"
	"testing"
)

func TestFastPowering(t *testing.T) {
	tests := []struct {
		name     string
		base     float64
		power    int
		expected float64
		epsilon  float64
	}{
		{
			name:     "Large Positive Base and Power",
			base:     5,
			power:    20,
			expected: 95367431640625,
			epsilon:  1e-6,
		},
		{
			name:     "Negative Base with Odd Power",
			base:     -3,
			power:    5,
			expected: -243,
			epsilon:  1e-6,
		},
		{
			name:     "Fractional Base with Positive Power",
			base:     0.5,
			power:    4,
			expected: 0.0625,
			epsilon:  1e-6,
		},
		{
			name:     "Very Large Power",
			base:     2,
			power:    1000000,
			expected: math.Inf(1),
			epsilon:  1e-6,
		},
		{
			name:     "Base 1 with Positive Power",
			base:     1,
			power:    100,
			expected: 1,
			epsilon:  1e-6,
		},
		{
			name:     "Base 1 with Negative Power",
			base:     1,
			power:    -50,
			expected: 1,
			epsilon:  1e-6,
		},
		{
			name:     "Base 1 with Zero Power",
			base:     1,
			power:    0,
			expected: 1,
			epsilon:  1e-6,
		},
		{
			name:     "Very Small Fractional Base with Large Power",
			base:     0.1,
			power:    100,
			expected: 1e-100,
			epsilon:  1e-106,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fastPowering(tt.base, tt.power)
			if math.Abs(result-tt.expected) > tt.epsilon {
				t.Errorf("fastPowering(%f, %d) = %g, expected %g", tt.base, tt.power, result, tt.expected)
			}
		})
	}
}
