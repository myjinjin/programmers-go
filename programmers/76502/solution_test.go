package programmers

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "[](){}",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "}]()[{",
			expected: 2,
		},
		{
			name:     "Example 3",
			s:        "[)(]",
			expected: 0,
		},
		{
			name:     "Example 4",
			s:        "}}}",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
