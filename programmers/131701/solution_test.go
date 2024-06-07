package programmers

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		elements []int
		expected int
	}{
		{
			name:     "Example 1",
			elements: []int{7, 9, 1, 1, 4},
			expected: 18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.elements)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
