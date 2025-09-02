package main

import (
	"fmt"
)

func maxWidthRamp(nums []int) int {
	maxWidth := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] <= nums[j] && i < j {
				fmt.Println(nums[i], nums[j], i, j)
				width := j - i
				if width > maxWidth {
					maxWidth = width
				}
			}
		}
	}
	return maxWidth
}

func main() {
	nums := []int{6, 0, 8, 2, 1, 5}
	fmt.Print(maxWidthRamp(nums))

}
