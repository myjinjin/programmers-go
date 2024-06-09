package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		arr1     [][]int
		arr2     [][]int
		expected [][]int
	}{
		{
			name:     "Example 1",
			arr1:     [][]int{{1, 4}, {3, 2}, {4, 1}},
			arr2:     [][]int{{3, 3}, {3, 3}},
			expected: [][]int{{15, 15}, {15, 15}, {15, 15}},
		},
		{
			name:     "Example 2",
			arr1:     [][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}},
			arr2:     [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}},
			expected: [][]int{{22, 22, 11}, {36, 28, 18}, {29, 20, 14}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
