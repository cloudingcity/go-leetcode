package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := New(numCourses)
	for _, prerequisite := range prerequisites {
		g.AddEdge(prerequisite[1], prerequisite[0])
	}
	return g.IsCycle()
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

func (g *Graph) IsCycle() bool {
	visited := make([]bool, g.vertices)
	recStack := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		if g.IsCycleUtil(i, visited, recStack) {
			return false
		}
	}
	return true
}

func (g *Graph) IsCycleUtil(v int, visited, recStack []bool) bool {
	if recStack[v] {
		return true
	}
	if visited[v] {
		return false
	}

	visited[v] = true
	recStack[v] = true
	for _, w := range g.adj[v] {
		if g.IsCycleUtil(w, visited, recStack) {
			return true
		}
	}
	recStack[v] = false
	return false
}
