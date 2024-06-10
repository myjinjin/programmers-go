package programmers

func solution(numbers []int64) []int64 {
	results := make([]int64, 0)
	for _, n := range numbers {
		rightMostZero := 1
		for int(n)&rightMostZero != 0 {
			rightMostZero <<= 1
		}

		result := (int(n) | rightMostZero)

		if int(n)&(rightMostZero>>1) != 0 {
			result &^= (rightMostZero >> 1)
		}

		results = append(results, int64(result))
	}
	return results
}
