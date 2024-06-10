package programmers

func solution(number string, k int) string {
	result := []byte{number[0]}

	for i := 1; i < len(number); i++ {
		curr := number[i]
		for len(result) > 0 && result[len(result)-1] < curr && k > 0 {
			result = result[:len(result)-1]
			k--
		}
		result = append(result, curr)
	}
	if len(result) > len(number)-k {
		result = result[:len(number)-k]
	}

	return string(result)
}
