package programmers

func solution(numbers []int, target int) int {
	return dfs(numbers, target, 0, 0)
}

func dfs(numbers []int, target int, index int, curr int) int {
	if index == len(numbers) {
		if curr == target {
			return 1
		}
		return 0
	}

	count := 0
	count += dfs(numbers, target, index+1, curr+numbers[index])
	count += dfs(numbers, target, index+1, curr-numbers[index])
	return count
}
