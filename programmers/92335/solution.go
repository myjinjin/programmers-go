package programmers

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(n int, k int) int {
	kbaseStr := toKBaseStr(n, k)
	candidates := extractCandidates(kbaseStr)
	count := 0
	for _, c := range candidates {
		if isPrime(c) {
			count++
		}
	}
	return count
}

func toKBaseStr(n int, k int) string {
	var sb strings.Builder
	for n != 0 {
		remainder := n % k
		sb.WriteString(fmt.Sprintf("%d", remainder))
		n /= k
	}

	runes := []rune(sb.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func extractCandidates(s string) []int {
	candidates := strings.Split(s, "0")
	result := make([]int, 0)
	for _, c := range candidates {
		if c != "" {
			candidateInt, _ := strconv.Atoi(c)
			result = append(result, candidateInt)
		}
	}
	return result
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
