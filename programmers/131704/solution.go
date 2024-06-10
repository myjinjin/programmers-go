package programmers

func solution(order []int) int {
	secondaryStack := make([]int, 0)
	containerBelt := make([]int, len(order))
	for i := range containerBelt {
		containerBelt[i] = i + 1
	}
	count := 0
	for len(order) > 0 {
		currOrder := order[0]

		if len(containerBelt) > 0 {
			if currOrder == containerBelt[0] {
				containerBelt = containerBelt[1:]
				count++
				order = order[1:]
			} else if len(secondaryStack) > 0 && currOrder == secondaryStack[len(secondaryStack)-1] {
				count++
				secondaryStack = secondaryStack[:len(secondaryStack)-1]
				order = order[1:]
			} else {
				secondaryStack = append(secondaryStack, containerBelt[0])
				containerBelt = containerBelt[1:]
			}
		} else if len(secondaryStack) > 0 {
			if currOrder == secondaryStack[len(secondaryStack)-1] {
				count++
				secondaryStack = secondaryStack[:len(secondaryStack)-1]
				order = order[1:]
			} else {
				break
			}
		}
	}

	return count
}
