package programmers

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int64
	}{
		{
			name:     "Example 1",
			n:        4,
			expected: 5,
		},
		{
			name:     "Example 2",
			n:        3,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
