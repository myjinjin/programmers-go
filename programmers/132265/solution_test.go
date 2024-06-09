package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		topping  []int
		expected int
	}{
		{
			name:     "Example 1",
			topping:  []int{1, 2, 1, 3, 1, 4, 1, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			topping:  []int{1, 2, 3, 1, 4},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.topping)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
