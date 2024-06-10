package programmers

func solution(queue1 []int, queue2 []int) int {
	oneSum := 0
	twoSum := 0
	n := len(queue1)

	for i := 0; i < n; i++ {
		oneSum += queue1[i]
		twoSum += queue2[i]
	}

	count := 0
	for oneSum != twoSum {
		if count > n*3 {
			return -1
		}
		if oneSum > twoSum {
			curr := queue1[0]
			queue1 = queue1[1:]
			queue2 = append(queue2, curr)
			count++
			oneSum -= curr
			twoSum += curr
		} else {
			curr := queue2[0]
			queue2 = queue2[1:]
			queue1 = append(queue1, curr)
			count++
			twoSum -= curr
			oneSum += curr
		}
		if len(queue1) == 0 || len(queue2) == 0 {
			return -1
		}
	}

	return count
}
