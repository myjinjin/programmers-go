package programmers

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{2, 6, 8, 14},
			expected: 168,
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 3},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
