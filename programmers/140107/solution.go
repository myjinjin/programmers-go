package programmers

import "math"

func solution(k int, d int) int64 {
	var count int64
	for i := 0; i*k <= d; i++ {
		x := i * k
		y := int(math.Sqrt(float64(d*d - x*x)))
		count += int64(y/k + 1)
	}
	return count
}
