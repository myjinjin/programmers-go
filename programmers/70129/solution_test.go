package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected []int
	}{
		{
			name:     "Example 1",
			s:        "110010101001",
			expected: []int{3, 8},
		},
		{
			name:     "Example 2",
			s:        "01110",
			expected: []int{3, 3},
		},
		{
			name:     "Example 3",
			s:        "1111111",
			expected: []int{4, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.s)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
