package main

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		edges    [][]int
		expected []int
	}{
		// Caso 1: Un ciclo con la última arista [2,3]
		{
			edges:    [][]int{{1, 2}, {1, 3}, {2, 3}},
			expected: []int{2, 3},
		},
		// Caso 2: El ciclo ocurre con la arista [1,4]
		{
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expected: []int{1, 4},
		},
		// Caso 3: Otro ciclo con la arista [4,6]
		{
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 4}},
			expected: []int{6, 4},
		},
		// Caso 4: Ciclo con la última arista [1,3]
		{
			edges:    [][]int{{1, 2}, {2, 3}, {3, 1}},
			expected: []int{3, 1},
		},
		// Caso 5: Ciclo más complejo con la última arista [5,6]
		{
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {2, 6}},
			expected: []int{2, 6},
		},
	}

	for _, test := range tests {
		result := findRedundantConnection(test.edges)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Para edges %v, esperado %v pero obtenido %v", test.edges, test.expected, result)
		}
	}
}
