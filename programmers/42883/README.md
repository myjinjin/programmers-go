# [큰 수 만들기](https://school.programmers.co.kr/learn/courses/30/lessons/42883)

## Problem

### 문제 설명

어떤 숫자에서 `k`개의 수를 제거했을 때 얻을 수 있는 가장 큰 숫자를 구하려 합니다.

예를 들어, 숫자 `1924`에서 수 두 개를 제거하면 `[19, 12, 14, 92, 94, 24]` 를 만들 수 있습니다. 이 중 가장 큰 숫자는 `94` 입니다.

문자열 형식으로 숫자 `number`와 제거할 수의 개수 `k`가 `solution` 함수의 매개변수로 주어집니다. `number`에서 `k` 개의 수를 제거했을 때 만들 수 있는 수 중 가장 큰 숫자를 문자열 형태로 return 하도록 solution 함수를 완성하세요.

### 제한 조건

- `number`는 `2`자리 이상, `1,000,000`자리 이하인 숫자입니다.
- `k`는 `1` 이상 `number`의 자릿수 미만인 자연수입니다.

### 입출력 예

| number | k | return |
| --- | --- | --- |
| "1924" | 2 | "94" |
| "1231234" | 3 | "3234" |
| "4177252841" | 4 | "775841" |

## Solution

```go
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
```