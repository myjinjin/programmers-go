# [행렬의 곱셈](https://school.programmers.co.kr/learn/courses/30/lessons/12949)

## Problem

### 문제 설명

2차원 행렬 `arr1`과 `arr2`를 입력받아, `arr1`에 `arr2`를 곱한 결과를 반환하는 함수, solution을 완성해주세요.

### 제한 조건

- 행렬 `arr1`, `arr2`의 행과 열의 길이는 `2` 이상 `100` 이하입니다.
- 행렬 `arr1`, `arr2`의 원소는 `-10` 이상 `20` 이하인 자연수입니다.
- 곱할 수 있는 배열만 주어집니다.

### 입출력 예

| arr1 | arr2 | return |
| --- | --- | --- |
| [[1, 4], [3, 2], [4, 1]] | [[3, 3], [3, 3]] | [[15, 15], [15, 15], [15, 15]] |
| [[2, 3, 2], [4, 2, 4], [3, 1, 4]] | [[5, 4, 3], [2, 4, 1], [3, 1, 1]] | [[22, 22, 11], [36, 28, 18], [29, 20, 14]] |

## Solution

```go
func solution(arr1 [][]int, arr2 [][]int) [][]int {
	sum := make([][]int, len(arr1))
	for i := range sum {
		sum[i] = make([]int, len(arr2[0]))
	}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2[0]); j++ {
			for k := 0; k < len(arr2); k++ {
				sum[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}
	return sum
}
```