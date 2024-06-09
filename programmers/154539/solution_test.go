package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "Example 1",
			numbers:  []int{2, 3, 3, 5},
			expected: []int{3, 5, 5, -1},
		},
		{
			name:     "Example 2",
			numbers:  []int{9, 1, 5, 3, 6, 2},
			expected: []int{-1, 5, 6, 6, -1, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.numbers)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
