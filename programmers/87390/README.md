# [n^2 배열 자르기](https://school.programmers.co.kr/learn/courses/30/lessons/87390)

## Problem

### 문제 설명

정수 `n`, `left`, `right`가 주어집니다. 다음 과정을 거쳐서 1차원 배열을 만들고자 합니다.

1. `n`행 `n`열 크기의 비어있는 2차원 배열을 만듭니다.
2. `i = 1, 2, 3, ..., n`에 대해서, 다음 과정을 반복합니다.
    - 1행 1열부터 `i`행 `i`열까지의 영역 내의 모든 빈 칸을 숫자 `i`로 채웁니다.
3. 1행, 2행, ..., `n`행을 잘라내어 모두 이어붙인 새로운 1차원 배열을 만듭니다.
4. 새로운 1차원 배열을 `arr`이라 할 때, `arr[left]`, `arr[left+1]`, ..., `arr[right]`만 남기고 나머지는 지웁니다.

정수 `n`, `left`, `right`가 매개변수로 주어집니다. 주어진 과정대로 만들어진 1차원 배열을 return 하도록 solution 함수를 완성해주세요.

### 제한사항

- 1 ≤ `n` ≤ 10^7
- 0 ≤ `left` ≤ `right` < n^2
- `right` - `left` < 10^5

### 입출력 예

| n | left | right | result |
| --- | --- | --- | --- |
| 3 | 2| 5| [3,2,2,3] |
| 4| 7| 14 | [4,3,3,3,4,4,4,4] |

## Solution

```go
func solution(n int, left int64, right int64) []int {
	result := make([]int, 0, right-left+1)
	for i := left; i < right+1; i++ {
		row := int(i) / n
		col := int(i) % n
		result = append(result, max(row, col)+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```