package main

import "slices"

func minimumArea(grid [][]int) int {
	yMin := 0
	for !rowContainsOne(grid, yMin) {
		yMin++
	}
	yMax := len(grid) - 1
	for !rowContainsOne(grid, yMax) {
		yMax--
	}
	xMin := 0
	for !columnContainsOne(grid, xMin) {
		xMin++
	}
	xMax := len(grid[0]) - 1
	for !columnContainsOne(grid, xMax) {
		xMax--
	}
	return (xMax - xMin + 1) * (yMax - yMin + 1)
}

func rowContainsOne(grid [][]int, row int) bool {
	return slices.Contains(grid[row], 1)
}

func columnContainsOne(grid [][]int, col int) bool {
	for y := 0; y < len(grid); y++ {
		if grid[y][col] == 1 {
			return true
		}
	}
	return false
}
