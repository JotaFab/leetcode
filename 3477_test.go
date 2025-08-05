package main

import "testing"

func TestNumOfUnplacedFruits(t *testing.T) {
	tests := []struct {
		name    string
		fruits  []int
		baskets []int
		want    int
	}{
		{
			name:    "All fruits can be placed",
			fruits:  []int{1, 2, 3},
			baskets: []int{3, 2, 1},
			want:    0,
		},
		{
			name:    "Some fruits cannot be placed",
			fruits:  []int{5, 3, 2, 1},
			baskets: []int{4, 2, 1, 1},
			want:    1,
		},
		{
			name:    "Empty baskets",
			fruits:  []int{1, 2, 3},
			baskets: []int{},
			want:    3,
		},
		{
			name:    "No matching baskets",
			fruits:  []int{5, 6, 7},
			baskets: []int{1, 2, 3},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfUnplacedFruits(tt.fruits, tt.baskets)
			if got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
