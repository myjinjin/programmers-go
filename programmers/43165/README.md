# [타겟 넘버](https://school.programmers.co.kr/learn/courses/30/lessons/43165)

## Problem

### 문제 설명

`n`개의 음이 아닌 정수들이 있습니다. 이 정수들을 순서를 바꾸지 않고 적절히 더하거나 빼서 타겟 넘버를 만들려고 합니다. 예를 들어 `[1, 1, 1, 1, 1]`로 숫자 `3`을 만들려면 다음 다섯 방법을 쓸 수 있습니다.

```
-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3
```

사용할 수 있는 숫자가 담긴 배열 `numbers`, 타겟 넘버 `target`이 매개변수로 주어질 때 숫자를 적절히 더하고 빼서 타겟 넘버를 만드는 방법의 수를 return 하도록 solution 함수를 작성해주세요.

### 제한사항

- 주어지는 숫자의 개수는 `2`개 이상 `20`개 이하입니다.
- 각 숫자는 `1` 이상 `50` 이하인 자연수입니다.
- 타겟 넘버는 `1` 이상 `1000` 이하인 자연수입니다.

### 입출력 예

| numbers | target | return |
| --- | --- | --- |
| [1, 1, 1, 1, 1] | 3 | 5 |
| [4, 1, 2, 1] | 4 | 2 |

## Solution

```go
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
```