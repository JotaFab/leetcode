package main

func countSquares(matrix [][]int) int {

	m, n, cnt := len(matrix), len(matrix[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i > 0 && j > 0 {
				matrix[i][j] = 1 + min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
			}
			cnt += matrix[i][j]
		}
	}
	return cnt
}
