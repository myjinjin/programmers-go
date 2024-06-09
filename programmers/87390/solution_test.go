package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		left     int64
		right    int64
		expected []int
	}{
		{
			name:     "Example 1",
			n:        3,
			left:     2,
			right:    5,
			expected: []int{3, 2, 2, 3},
		},
		{
			name:     "Example 2",
			n:        4,
			left:     7,
			right:    14,
			expected: []int{4, 3, 3, 3, 4, 4, 4, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.n, tc.left, tc.right)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
