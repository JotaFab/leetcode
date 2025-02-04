package main

import "testing"

func TestIsSpecialArray(t *testing.T) {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1}, true},                  // Single element, always true
		{[]int{2, 1, 4}, true},            // Alternating even-odd
		{[]int{4, 3, 1, 6}, false},        // Two consecutive odd numbers
		{[]int{1, 2, 3, 4}, true},         // Alternating pattern
		{[]int{5, 7, 9}, false},           // All odd, not special
		{[]int{2, 4, 6, 8}, false},        // All even, not special
		{[]int{1, 2, 3, 4, 5, 6}, true},   // Long alternating pattern
		{[]int{10, 21, 32, 43, 54}, true}, // Alternating even-odd with large numbers
		{[]int{0, 1, 2, 3, 4, 5}, true},   // Includes zero, still alternating
		{[]int{1, 1, 2, 2}, false},
	}
	for _, tc := range testCases {
		result := isArraySpecial(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
		}
	}
}
