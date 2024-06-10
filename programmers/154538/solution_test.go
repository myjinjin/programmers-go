package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			x:        10,
			y:        40,
			n:        5,
			expected: 2,
		},
		{
			name:     "Example 2",
			x:        10,
			y:        40,
			n:        30,
			expected: 1,
		},
		{
			name:     "Example 3",
			x:        2,
			y:        5,
			n:        4,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.x, tc.y, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
