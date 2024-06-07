package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		k        int
		d        int
		expected int64
	}{
		{
			name:     "Example 1",
			k:        2,
			d:        4,
			expected: 6,
		},
		{
			name:     "Example 2",
			k:        1,
			d:        5,
			expected: 26,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.k, tc.d)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
