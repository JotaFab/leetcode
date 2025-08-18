package main

import "sort"

func numOfUnplacedFruitsIII(fruits []int, baskets []int) int {
	n := len(fruits)
	bs := make([]int, n)
	copy(bs, baskets)

	sort.Ints(bs)

	ans := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; i++ {
			if fruits[i] <= baskets[j] {
				baskets[j] = 0
				ans--
				break
			}
		}
	}
	return ans
}
