package programmers

import (
	"sort"
)

func solution(k int, tangerine []int) int {
	tangerineCount := make(map[int]int)
	for _, t := range tangerine {
		tangerineCount[t]++
	}

	counts := make([]int, 0, len(tangerineCount))
	for _, count := range tangerineCount {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	result := 0
	for _, count := range counts {
		k -= count
		result++
		if k <= 0 {
			break
		}
	}
	return result
}
