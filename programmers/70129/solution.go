package programmers

import (
	"fmt"
	"strings"
)

func solution(s string) []int {
	convertCount := 0
	deleteZeroCount := 0
	for s != "1" {
		convertCount++
		zeroCount := strings.Count(s, "0")
		deleteZeroCount += zeroCount
		oneCount := len(s) - zeroCount
		s = fmt.Sprintf("%b", oneCount)
	}

	return []int{convertCount, deleteZeroCount}
}
