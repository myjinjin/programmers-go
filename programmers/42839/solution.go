package programmers

import (
	"strconv"
)

func solution(numbers string) int {
	numSubsets := generateSubsets(numbers)
	numMap := make(map[int]bool)
	for _, numStr := range numSubsets {
		n, _ := strconv.Atoi(numStr)
		if isPrime(n) {
			numMap[n] = true
		}
	}

	return len(numMap)
}

func generateSubsets(str string) []string {
	var subsets []string
	n := len(str)
	for i := 1; i <= n; i++ {
		subsets = append(subsets, permuteHelper(str, i)...)
	}
	return subsets
}

func permuteHelper(str string, length int) []string {
	var result []string
	if length == 1 {
		for _, char := range str {
			result = append(result, string(char))
		}
		return result
	}

	for i := 0; i < len(str); i++ {
		char := string(str[i])
		remaining := str[:i] + str[i+1:]
		for _, subPerm := range permuteHelper(remaining, length-1) {
			result = append(result, char+subPerm)
		}
	}
	return result
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
