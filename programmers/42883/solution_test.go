package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		number   string
		k        int
		expected string
	}{
		{
			name:     "Example 1",
			number:   "1924",
			k:        2,
			expected: "94",
		},
		{
			name:     "Example 2",
			number:   "1231234",
			k:        3,
			expected: "3234",
		},
		{
			name:     "Example 3",
			number:   "4177252841",
			k:        4,
			expected: "775841",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.number, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
