# [숫자 변환하기](https://school.programmers.co.kr/learn/courses/30/lessons/154538)

## Problem

### 문제 설명

자연수 `x`를 `y`로 변환하려고 합니다. 사용할 수 있는 연산은 다음과 같습니다.

- `x`에 `n`을 더합니다
- `x`에 `2`를 곱합니다.
- `x`에 `3`을 곱합니다.

자연수 `x`, `y`, `n`이 매개변수로 주어질 때, `x`를 `y`로 변환하기 위해 필요한 최소 연산 횟수를 return하도록 solution 함수를 완성해주세요. 이때 `x`를 `y`로 만들 수 없다면 `-1`을 return 해주세요.

### 제한사항

- `1 ≤ x ≤ y ≤ 1,000,000`
- `1 ≤ n < y`

### 입출력 예

| x | y | n | result |
| -- | -- | -- | -- |
| 10 | 40 | 5 | 2 |
| 10 | 40 | 30 | 1 |
| 2 | 5 | 4 | -1 |

## Solution

```go
import "math"

func solution(x int, y int, n int) int {
	dp := make([]int, y+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[x] = 0

	for i := x; i <= y; i++ {
		if i+n <= y && dp[i+n] > dp[i]+1 {
			dp[i+n] = dp[i] + 1
		}
		if i*2 <= y && dp[i*2] > dp[i]+1 {
			dp[i*2] = dp[i] + 1
		}
		if i*3 <= y && dp[i*3] > dp[i]+1 {
			dp[i*3] = dp[i] + 1
		}
	}

	if dp[y] == math.MaxInt32 {
		return -1
	}
	return dp[y]
}
```