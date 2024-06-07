package programmers

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		want     []string
		number   []int
		discount []string
		expected int
	}{
		{
			name:     "Example 1",
			want:     []string{"banana", "apple", "rice", "pork", "pot"},
			number:   []int{3, 2, 2, 2, 1},
			discount: []string{"chicken", "apple", "apple", "banana", "rice", "apple", "pork", "banana", "pork", "rice", "pot", "banana", "apple", "banana"},
			expected: 3,
		},
		{
			name:     "Example 2",
			want:     []string{"apple"},
			number:   []int{10},
			discount: []string{"banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana"},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.want, tc.number, tc.discount)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
