package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		plans    [][]string
		expected []string
	}{
		{
			name:     "Example 1",
			plans:    [][]string{{"korean", "11:40", "30"}, {"english", "12:10", "20"}, {"math", "12:30", "40"}},
			expected: []string{"korean", "english", "math"},
		},
		{
			name:     "Example 2",
			plans:    [][]string{{"science", "12:40", "50"}, {"music", "12:20", "40"}, {"history", "14:00", "30"}, {"computer", "12:30", "100"}},
			expected: []string{"science", "history", "computer", "music"},
		},
		{
			name:     "Example 3",
			plans:    [][]string{{"aaa", "12:00", "20"}, {"bbb", "12:10", "30"}, {"ccc", "12:40", "10"}},
			expected: []string{"bbb", "ccc", "aaa"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.plans)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
