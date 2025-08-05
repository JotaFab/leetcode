package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	cnt := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fruits[i] <= baskets[j] {
				baskets[j] = 0
				cnt--
				break
			}
		}
	}
	return cnt
}
