package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		order    []int
		expected int
	}{
		{
			name:     "Example 1",
			order:    []int{4, 3, 1, 2, 5},
			expected: 2,
		},
		{
			name:     "Example 2",
			order:    []int{5, 4, 3, 2, 1},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.order)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
