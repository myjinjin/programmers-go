package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  string
		expected int
	}{
		{
			name:     "Example 1",
			numbers:  "17",
			expected: 3,
		},
		{
			name:     "Example 2",
			numbers:  "011",
			expected: 2,
		},
		{
			name:     "Example 3",
			numbers:  "121", // 2, 11, 211
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.numbers)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
