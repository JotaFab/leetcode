package main

func findRedundantConnection(edges [][]int) []int {
	dsu := NewDSU(len(edges))
	for _, edge := range edges {
		if !dsu.Union(edge[0], edge[1]) {
			return edge
		}

	}
	return nil
}

type DSU struct {
	size, rep []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		size: make([]int, n+1),
		rep:  make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		dsu.rep[i] = i
		dsu.size[i] = 1
	}
	return dsu
}
func (dsu *DSU) Find(x int) int {
	if dsu.rep[x] != x {
		dsu.rep[x] = dsu.Find(dsu.rep[x]) // Path compression
	}
	return dsu.rep[x]
}
func (dsu *DSU) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)

	if rootX == rootY {
		return false // Cycle detected
	}
	if dsu.size[rootX] > dsu.size[rootY] {
		dsu.rep[rootY] = rootX
	} else if dsu.size[rootX] == dsu.size[rootY] {
		dsu.rep[rootX] = rootY
	} else {
		dsu.rep[rootY] = rootX
		dsu.size[rootX]++
	}
	return true
}
