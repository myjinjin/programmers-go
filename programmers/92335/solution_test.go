package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        437674,
			k:        3,
			expected: 3,
		},
		{
			name:     "Example 2",
			n:        110011,
			k:        10,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
