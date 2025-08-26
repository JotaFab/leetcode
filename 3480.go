package main

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	total := int64(n) * int64(n+1) / 2
	best := int64(0)

    for i, _ := range conflictingPairs {
        temp := make([][2]int, 0)
		for j, p := range conflictingPairs {
			if i == j {
				continue
			}
			a, b := p[0], p[1]
			if a > b {
				a, b = b, a
			}
			temp = append(temp, [2]int{a, b})

		}
		if len(temp) == 0 {
			best = max(best, total)
			continue
		}
		L, R := temp[0][0], temp[0][1]
		for _, iv := range temp {
			if iv[0] < L {
				L = iv[0]
			}
			if iv[1] > R {
				R = iv[1]
			}
		}
		invalid := int64(L) * int64(n-R+1)
		valid := total - invalid
		if valid > best {
			best = valid
		}
	}
	return best
    
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
