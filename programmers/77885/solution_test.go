package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int64
		expected []int64
	}{
		{
			name:     "Example 1",
			numbers:  []int64{2, 7},
			expected: []int64{3, 11},
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
