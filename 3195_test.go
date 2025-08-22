package main

import "testing"

func TestMinimumArea(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "simple square",
			grid: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			want: 4, // sub-rectángulo de 2x2
		},
		{
			name: "single one",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 1, // un solo pixel
		},
		{
			name: "line horizontal",
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
			want: 3, // ancho de 3, alto de 1
		},
		{
			name: "line vertical",
			grid: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			want: 3, // ancho de 1, alto de 3
		},
		{
			name: "ones in corners",
			grid: [][]int{
				{1, 0, 0, 1},
				{0, 0, 0, 0},
				{1, 0, 0, 1},
			},
			want: 12, // rectángulo completo que cubre (0,0) hasta (3,2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumArea(tt.grid)
			if got != tt.want {
				t.Errorf("minimumArea() = %d, want %d", got, tt.want)
			}
		})
	}
}
