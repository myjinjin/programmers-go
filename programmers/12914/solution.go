package programmers

func solution(n int) int64 {
	dp := make([]int64, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1234567
	}
	return dp[n]
}
