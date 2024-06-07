package programmers

func solution(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if a <= b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}
	return gcd(b, (a % b))
}
