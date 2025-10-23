package main

import (
	"log"
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)

	// frequency map for equal counts
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ans := 1
	n := len(freq)
	

		

	
	

	return ans
}



func main() {
	log.Println(maxFrequency([]int{1, 90}, 1, 76)) // 1
	log.Println(maxFrequency([]int{88, 53}, 27, 2)) // 2
}
