package main

import "sort"

func numberOfPairs(points [][]int) int {
	var ans int
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	n := len(points)
	for i, A := range points {
		for j := i + 1; j < n; j++ {
			B := points[j]

			if A[1] < B[1] {
				continue
			}
			valid := true
			for k := i + 1; k < j; k++ {
				P := points[k]
				if B[1] <= P[1] && P[1] <= A[1] {
					valid = false
					break
				}
			}
			if valid {
				ans++
			}
		}
	}

	return ans
}
