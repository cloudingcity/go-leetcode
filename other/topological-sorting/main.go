package main

// see https://www.geeksforgeeks.org/topological-sorting/
type Graph struct {
	vertices int
	adj      [][]int
}

func New(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adj:      make([][]int, vertices),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
}

func (g *Graph) topologicalSort() []int {
	var stack []int
	visited := make([]bool, g.vertices)
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, &stack)
		}
	}

	var res []int
	for i := g.vertices - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return res
}

func (g *Graph) topologicalSortUtil(v int, visited []bool, stack *[]int) {
	visited[v] = true
	for _, w := range g.adj[v] {
		if !visited[w] {
			g.topologicalSortUtil(w, visited, stack)
		}
	}
	*stack = append(*stack, v)
}
