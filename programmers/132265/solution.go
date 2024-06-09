package programmers

func solution(topping []int) int {
	n := len(topping)
	left := make([]int, n) // 왼쪽~i번째까지 토핑 종류 개수
	leftTopping := make(map[int]int)
	for i := 0; i < n; i++ {
		leftTopping[topping[i]]++
		left[i] = len(leftTopping)
	}
	right := make([]int, n) // i번째~오른쪽까지 토핑 종류 개수
	rightTopping := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		rightTopping[topping[i]]++
		right[i] = len(rightTopping)
	}
	count := 0
	for i := 1; i < n; i++ {
		if right[i] == left[i-1] {
			count++
		}
	}
	return count
}
