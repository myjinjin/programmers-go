package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name      string
		k         int
		tangerine []int
		expected  int
	}{
		{
			name:      "Example 1",
			k:         6,
			tangerine: []int{1, 3, 2, 5, 4, 5, 2, 3},
			expected:  3,
		},
		{
			name:      "Example 2",
			k:         4,
			tangerine: []int{1, 3, 2, 5, 4, 5, 2, 3},
			expected:  2,
		},
		{
			name:      "Example 3",
			k:         2,
			tangerine: []int{1, 1, 1, 1, 2, 2, 2, 3},
			expected:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.k, tc.tangerine)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
