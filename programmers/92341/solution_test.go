package programmers

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		fees     []int
		records  []string
		expected []int
	}{
		{
			name:     "Example 1",
			fees:     []int{180, 5000, 10, 600},
			records:  []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"},
			expected: []int{14600, 34400, 5000},
		},
		{
			name:     "Example 2",
			fees:     []int{120, 0, 60, 591},
			records:  []string{"16:00 3961 IN", "16:00 0202 IN", "18:00 3961 OUT", "18:00 0202 OUT", "23:58 3961 IN"},
			expected: []int{0, 591},
		},
		{
			name:     "Example 3",
			fees:     []int{1, 461, 1, 10},
			records:  []string{"00:00 1234 IN"},
			expected: []int{14841},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.fees, tc.records)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
