package programmers

import "math"

func solution(x int, y int, n int) int {
	dp := make([]int, y+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[x] = 0

	for i := x; i <= y; i++ {
		if i+n <= y && dp[i+n] > dp[i]+1 {
			dp[i+n] = dp[i] + 1
		}
		if i*2 <= y && dp[i*2] > dp[i]+1 {
			dp[i*2] = dp[i] + 1
		}
		if i*3 <= y && dp[i*3] > dp[i]+1 {
			dp[i*3] = dp[i] + 1
		}
	}

	if dp[y] == math.MaxInt32 {
		return -1
	}
	return dp[y]
}
