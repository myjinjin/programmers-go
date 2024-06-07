package programmers

import "reflect"

func solution(want []string, number []int, discount []string) int {
	availableRegisterDay := 0
	windowSize := 10

	wantCount := make(map[string]int)
	for i := 0; i < len(want); i++ {
		wantCount[want[i]] = number[i]
	}

	discountCount := make(map[string]int)
	for start := 0; start < windowSize; start++ {
		discountCount[discount[start]]++
	}
	if reflect.DeepEqual(wantCount, discountCount) {
		availableRegisterDay++
	}
	for start := 1; start+windowSize-1 < len(discount); start++ {
		discountCount[discount[start-1]]--
		if discountCount[discount[start-1]] == 0 {
			delete(discountCount, discount[start-1])
		}
		discountCount[discount[start+windowSize-1]]++

		if reflect.DeepEqual(wantCount, discountCount) {
			availableRegisterDay++
		}
	}

	return availableRegisterDay
}
