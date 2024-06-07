package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		a        int
		b        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        8,
			a:        4,
			b:        7,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.n, tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
