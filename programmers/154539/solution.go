package programmers

func solution(numbers []int) []int {
	n := len(numbers)
	bigNum := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && numbers[stack[len(stack)-1]] < numbers[i] {
			bigNum[stack[len(stack)-1]] = numbers[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		bigNum[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	return bigNum
}
