package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := New(numCourses)
	for _, prerequisite := range prerequisites {
		g.AddEdge(prerequisite[0], prerequisite[1])
	}
	return g.FindOrder()
}

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

func (g *Graph) FindOrder() []int {
	var res []int
	visited := make([]bool, g.vertices)
	recStack := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		if visited[i] {
			continue
		}
		if g.IsCycleUtil(i, &res, visited, recStack) {
			return nil
		}
	}
	return res
}

func (g *Graph) IsCycleUtil(v int, res *[]int, visited, recStack []bool) bool {
	if recStack[v] {
		return true
	}
	if visited[v] {
		return false
	}

	visited[v] = true
	recStack[v] = true
	for _, w := range g.adj[v] {
		if g.IsCycleUtil(w, res, visited, recStack) {
			return true
		}
	}
	recStack[v] = false
	*res = append(*res, v)
	return false
}
