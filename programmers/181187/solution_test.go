package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		r1       int
		r2       int
		expected int64
	}{
		{
			name:     "Example 1",
			r1:       2,
			r2:       3,
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.r1, tc.r2)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
