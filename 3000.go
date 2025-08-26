package main


func areaOfMaxDiagonal(dimensions [][]int) int {
    maxD := -1
	maxArea := -1

    for _, r:= range dimensions {
        l, w := r[0], r[1]
		d := l*l + w*w
		area := l * w
		if d > maxD || (d == maxD && area > maxArea) {
			maxD = d
			maxArea = area
		}
    }
    return maxArea

}