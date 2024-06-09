package programmers

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
