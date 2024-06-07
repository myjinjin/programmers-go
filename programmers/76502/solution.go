package programmers

func solution(s string) int {
	count := 0

	for start := 0; start < len(s); start++ {
		rotated := s[start:] + s[:start]
		stack := make([]byte, 0, len(rotated))
		isValid := true
		for i := 0; i < len(rotated); i++ {
			curr := rotated[i]
			if isOpenBracket(curr) {
				stack = append(stack, curr)
			} else {
				if len(stack) == 0 || getMatchingBracket(curr) != stack[len(stack)-1] {
					isValid = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		if isValid && len(stack) == 0 {
			count++
		}
	}

	return count
}

func isOpenBracket(bracket byte) bool {
	switch bracket {
	case '[', '(', '{':
		return true
	}
	return false
}

func getMatchingBracket(closeBracket byte) byte {
	switch closeBracket {
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	}
	return 0
}
