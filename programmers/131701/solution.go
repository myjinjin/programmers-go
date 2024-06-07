package programmers

func solution(elements []int) int {
	uniqueSums := make(map[int]bool)
	calculateUniqueSums(elements, uniqueSums)
	return len(uniqueSums)
}

func calculateUniqueSums(elements []int, uniqueSums map[int]bool) {
	for start := 0; start < len(elements); start++ {
		rotated := append(elements[start:], elements[:start]...)
		dp := make([]int, len(rotated))
		dp[0] = rotated[0]
		uniqueSums[dp[0]] = true
		for i := 1; i < len(rotated); i++ {
			dp[i] = dp[i-1] + rotated[i]
			uniqueSums[dp[i]] = true
		}
	}
}
