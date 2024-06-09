package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			numbers:  []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5,
		},
		{
			name:     "Example 2",
			numbers:  []int{4, 1, 2, 1},
			target:   4,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.numbers, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
