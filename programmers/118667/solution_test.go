package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		queue1   []int
		queue2   []int
		expected int
	}{
		{
			name:     "Example 1",
			queue1:   []int{3, 2, 7, 2},
			queue2:   []int{4, 6, 5, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			queue1:   []int{1, 2, 1, 2},
			queue2:   []int{1, 10, 1, 2},
			expected: 7,
		},
		{
			name:     "Example 3",
			queue1:   []int{1, 1},
			queue2:   []int{1, 5},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.queue1, tc.queue2)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
