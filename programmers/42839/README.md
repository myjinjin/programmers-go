# [소수 찾기](https://school.programmers.co.kr/learn/courses/30/lessons/42839)

## Problem

### 문제 설명

한자리 숫자가 적힌 종이 조각이 흩어져있습니다. 흩어진 종이 조각을 붙여 소수를 몇 개 만들 수 있는지 알아내려 합니다.

각 종이 조각에 적힌 숫자가 적힌 문자열 `numbers`가 주어졌을 때, 종이 조각으로 만들 수 있는 소수가 몇 개인지 return 하도록 solution 함수를 완성해주세요.

### 제한사항

- `numbers`는 길이 `1` 이상 `7` 이하인 문자열입니다.
- `numbers`는 `0~9`까지 숫자만으로 이루어져 있습니다.
- `"013"`은 `0`, `1`, `3` 숫자가 적힌 종이 조각이 흩어져있다는 의미입니다.

### 입출력 예

| numbers | return |
| ---- | ---- |
| "17" | 3 |
| "011" | 2 |

## Solution

```go
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
```