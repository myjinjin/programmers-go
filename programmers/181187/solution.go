package programmers

import "math"

func solution(r1 int, r2 int) int64 {
	var count int64

	count += int64(r2-r1+1) * 4

	for x := 1; x <= r2; x++ {
		yMax := int(math.Sqrt(float64(r2*r2 - x*x)))
		yMin := 0
		if x <= r1 {
			yMin = int(math.Ceil(math.Sqrt(float64(r1*r1 - x*x))))
		}
		if yMin != 0 {
			count += int64(yMax-yMin+1) * 4
		} else {
			count += int64(yMax-yMin) * 4
		}
	}
	return count
}
