package programmers

func solution(n int, a int, b int) int {
	count := 0

	for a != b {
		count++
		a = (a + 1) / 2
		b = (b + 1) / 2
	}

	return count
}
