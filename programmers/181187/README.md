# [두 원 사이의 정수 쌍](https://school.programmers.co.kr/learn/courses/30/lessons/181187)

## Problem

### 문제 설명

x축과 y축으로 이루어진 2차원 직교 좌표계에 중심이 원점인 서로 다른 크기의 원이 두 개 주어집니다. 반지름을 나타내는 두 정수 `r1`, `r2`가 매개변수로 주어질 때, 두 원 사이의 공간에 x좌표와 y좌표가 모두 정수인 점의 개수를 return하도록 solution 함수를 완성해주세요.
※ 각 원 위의 점도 포함하여 셉니다.

### 제한 사항

- `1 ≤ r1 < r2 ≤ 1,000,000`

### 입출력 예

| r1 | r2 | result |
| -- | -- | ------ |
| 2 | 3 | 20 |

## Solution

```go
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
```