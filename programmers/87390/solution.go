package programmers

func solution(n int, left int64, right int64) []int {
	result := make([]int, 0, right-left+1)
	for i := left; i < right+1; i++ {
		row := int(i) / n
		col := int(i) % n
		result = append(result, max(row, col)+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
